# Frontend System Patterns: PowerPoint Translator App

## 1. Frontend Architecture Overview

The frontend architecture follows a structured approach with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                   Next.js 14 App Router                    │
│                                                             │
│  ┌─────────────────────┐      ┌─────────────────────────┐  │
│  │  Server Components  │      │    Client Components    │  │
│  │  ─────────────────  │      │  ─────────────────────  │  │
│  │  - Data fetching    │      │  - Interactive UI       │  │
│  │  - Session routing  │      │  - Slide rendering      │  │
│  │  - API integration  │      │  - Edit interfaces      │  │
│  └─────────────────────┘      └─────────────────────────┘  │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                Zustand State Stores                  │   │
│  │ ┌─────────────┐ ┌─────────────┐ ┌─────────────────┐ │   │
│  │ │sessionStore │ │ slideStore  │ │  commentStore   │ │   │
│  │ └─────────────┘ └─────────────┘ └─────────────────┘ │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                 Service Layer                        │   │
│  │ ┌─────────────┐ ┌─────────────┐ ┌─────────────────┐ │   │
│  │ │SessionSvc   │ │ PPTXSvc     │ │  CommentSvc     │ │   │
│  │ └─────────────┘ └─────────────┘ └─────────────────┘ │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
           │                  │                  │
           ▼                  ▼                  ▼
┌────────────────┐  ┌────────────────┐  ┌────────────────────┐
│   Supabase     │  │   pptxgenjs    │  │   Microservices    │
│  Auth/Storage  │  │  PPTX Library  │  │   REST APIs        │
└────────────────┘  └────────────────┘  └────────────────────┘
```

## 2. Component Structure & Organization

### 2.1. Directory Structure

```
/frontend
  /app                       # Next.js App Router
    /layout.tsx              # Root layout with providers
    /page.tsx                # Landing/login page
    /sessions                # Session routes
      /page.tsx              # Dashboard (Server Component)
      /new                   # New session route
        /page.tsx
      /[sessionId]           # Dynamic session routes
        /layout.tsx          # Session shell layout
        /page.tsx            # Overview
        /slides              # Slides routes
          /[slideNumber]     # Individual slide route
            /page.tsx        # SlideEditor page (Client)
        /share               # Sharing page
          /page.tsx
        /history             # History/audit page
          /page.tsx
        /export              # Export page
          /page.tsx
  
  /components                # React components
    /ui                      # shadcn/ui components
    /session                 # Session-related components
      SessionCard.tsx
      UploadWizard.tsx
      ShareModal.tsx
    /slide                   # Slide-related components
      SlideCanvas.tsx
      SlideNavigator.tsx
      TextChunkEditor.tsx
      MergePanel.tsx
      ReadingOrderOverlay.tsx
    /comments                # Comment-related components
      CommentThread.tsx
      CommentForm.tsx
      NotificationBadge.tsx
    /shared                  # Shared components
      LoadingSpinner.tsx
      ErrorBoundary.tsx

  /hooks                     # Custom React hooks
    /useSessionStore.ts      # Zustand store hooks
    /useSlideStore.ts
    /useCommentStore.ts
    /useNotificationStore.ts
    /useRealtimeSubscription.ts
    /usePPTXParser.ts
    
  /lib                       # Utilities and services
    /services                # Service layer
      SessionService.ts
      PPTXService.ts
      CommentService.ts
      MergeService.ts
      OrderService.ts
      AuditService.ts
      ExportService.ts
    /utils                   # Utility functions
      supabaseClient.ts
      pptxHelpers.ts
      formatters.ts
      validators.ts
  
  /styles                    # Global styles
    /globals.css             # Tailwind imports
  
  /types                     # TypeScript type definitions
    /models.ts               # Domain models
    /api.ts                  # API responses
    /supabase.ts             # Supabase types
```

## 3. Design Patterns & Architecture Decisions

### 3.1. State Management with Zustand

The frontend uses a **slice pattern** with Zustand for global state management:

```typescript
// sessionStore.ts
interface SessionState {
  currentSession: Session | null;
  role: 'owner' | 'reviewer';
  shareToken: string | null;
  
  setSession: (session: Session) => void;
  setRole: (role: 'owner' | 'reviewer') => void;
  setShareToken: (token: string) => void;
  // ...other actions
}

export const useSessionStore = create<SessionState>()(
  (set) => ({
    currentSession: null,
    role: 'owner',
    shareToken: null,
    
    setSession: (session) => set({ currentSession: session }),
    setRole: (role) => set({ role }),
    setShareToken: (token) => set({ shareToken: token }),
    // ...other actions
  })
);
```

Each store follows similar patterns:
- **Single source of truth** for its domain
- **Immutable updates** via Zustand's `set` function
- **Selector hooks** to prevent unnecessary re-renders
- **Action methods** that encapsulate logic

### 3.2. Service Layer Pattern

All external communication happens through service modules:

```typescript
// SessionService.ts
export class SessionService {
  private supabase: SupabaseClient;
  
  constructor() {
    this.supabase = createClientComponentClient();
  }
  
  async createSession(name: string, file: File): Promise<Session> {
    // Upload to Supabase Storage
    // Insert metadata to sessions table
    // Return new session
  }
  
  async generateShareLink(sessionId: string, permissions: SharePermissions): Promise<string> {
    // Create share token
    // Insert to session_shares table
    // Return shareable URL
  }
  
  // ...other methods
}
```

Benefits of this approach:
- **Encapsulation** of API/database logic
- **Reusability** across components
- **Testability** via mocking
- **Single responsibility** per service

### 3.3. Component Composition Pattern

Components are composed using a hierarchy that promotes reuse:

```jsx
// SlideEditorPage.tsx
export default function SlideEditorPage({ params }) {
  // Fetch data via Server Component patterns
  return (
    <div className="slide-editor-layout">
      <Sidebar>
        <SlideNavigator slides={slides} currentSlide={params.slideNumber} />
      </Sidebar>
      <MainContent>
        <SlideCanvas 
          slide={currentSlide} 
          editable={permissions.canEdit} 
        />
        {permissions.canEdit && (
          <>
            <TextChunkEditor />
            <MergePanel />
            <ReadingOrderOverlay />
          </>
        )}
      </MainContent>
      <RightPanel>
        <CommentThread slideId={params.slideNumber} />
      </RightPanel>
    </div>
  );
}
```

This composition pattern allows:
- **Conditional rendering** based on permissions
- **Flexible layouts** with sidebar, main content, and panels
- **Clear separation** of component responsibilities
- **Prop drilling minimization** through context or stores

### 3.4. Custom Hook Pattern

Custom hooks extract and reuse stateful logic:

```typescript
// usePPTXParser.ts
export function usePPTXParser() {
  const [parsing, setParsing] = useState(false);
  const [progress, setProgress] = useState(0);
  const [error, setError] = useState<Error | null>(null);
  
  const parseFile = useCallback(async (file: File) => {
    setParsing(true);
    setProgress(0);
    setError(null);
    
    try {
      // Load pptxgenjs dynamically
      const pptxgen = await import('pptxgenjs');
      
      // Parse the file
      // Update progress as slides are processed
      // Return the extracted metadata
      
      setProgress(100);
      return metadata;
    } catch (err) {
      setError(err as Error);
      return null;
    } finally {
      setParsing(false);
    }
  }, []);
  
  return { parseFile, parsing, progress, error };
}
```

Benefits:
- **Reusable logic** across components
- **Consistent state handling** for async operations
- **Encapsulated implementation details**
- **Testable** in isolation from UI

### 3.5. Server/Client Component Pattern

Following Next.js 14's App Router architecture:

- **Server Components** for:
  - Data fetching directly from Supabase
  - Initial session loading
  - Static UI elements
  - SEO and metadata

- **Client Components** for:
  - Interactive UI elements
  - State-driven interfaces
  - Event handling
  - Realtime subscriptions

## 4. Data Flow Patterns

### 4.1. PPTX Processing Flow

```
┌─────────────┐    ┌────────────┐    ┌────────────────┐    ┌────────────┐
│ File Upload │ -> │ PPTX Parse │ -> │ Metadata Store │ -> │ DB Storage │
└─────────────┘    └────────────┘    └────────────────┘    └────────────┘
                                          |
                                          v
┌──────────────┐    ┌───────────────┐    ┌────────────────┐
│ Export PPTX  │ <- │ PPTX Generate │ <- │ Edited Metadata│
└──────────────┘    └───────────────┘    └────────────────┘
```

### 4.2. Editing & Collaboration Flow

```
┌─────────────┐    ┌──────────────┐    ┌─────────────┐    ┌───────────┐
│ Text Update │ -> │ Local Buffer │ -> │ Save Action │ -> │ DB Update │
└─────────────┘    └──────────────┘    └─────────────┘    └───────────┘
                                                               |
                                                               v
┌───────────────┐    ┌───────────────────┐    ┌─────────────────────┐
│ Audit Log     │ <- │ Notification Event│ <- │ Supabase LISTEN     │
└───────────────┘    └───────────────────┘    └─────────────────────┘
```

### 4.3. Authentication Flow

```
┌───────────────┐    ┌──────────────┐    ┌─────────────────┐
│ Email/Password│ -> │ Supabase Auth│ -> │ JWT in LocalStorage│
└───────────────┘    └──────────────┘    └─────────────────┘
                                               |
                                               v
┌───────────────┐    ┌──────────────────┐    ┌─────────────────┐
│ API Requests  │ -> │ Authorization Header│ -> │ Supabase RLS    │
└───────────────┘    └──────────────────┘    └─────────────────┘
```

## 5. Frontend-Specific Design Patterns

### 5.1. PPTX Rendering Strategy
- **Canvas-based Rendering**: Using pptxgenjs to render slides to canvas elements
- **Overlay System**: Transparent div overlays for interactive elements
- **Position Mapping**: 1:1 mapping between PowerPoint coordinates and rendered web elements

### 5.2. State Persistence Strategy
- **Optimistic Updates**: UI updates immediately, then confirms with backend
- **Debounced Saves**: Text edits batch-saved after typing pauses
- **Offline Support**: LocalStorage for unsaved changes
- **Change Tracking**: Dirty state indicators for unsaved work

### 5.3. UI Component Patterns
- **Compound Components**: Related UI elements grouped as cohesive units
- **Render Props**: For complex, customizable UI elements
- **Controlled Components**: For form inputs and editable content
- **Portal Pattern**: For modals, tooltips, and floating editors

## 6. Frontend Communication Patterns

### 6.1. Microservice Communication
- **Service Modules**: Abstract HTTP calls behind clean interfaces
- **Error Handling**: Consistent error parsing and display
- **Request Queueing**: For offline support or rate limiting
- **Response Caching**: For frequently accessed data

### 6.2. Realtime Updates
- **Supabase Subscriptions**: For comments and notifications
- **UI Badge Updates**: Real-time counters for unread items
- **Notification Queue**: Manage incoming notifications with priority

### 6.3. Caching Strategy
- **React Query/SWR**: For data that needs caching beyond Zustand (optional)
- **LocalStorage**: For user preferences and session data
- **Memory Cache**: For rendered slide data to improve performance

## 7. Performance Patterns

### 7.1. Rendering Optimizations
- **Virtualization**: For slide thumbnails and large comment threads
- **Code Splitting**: Based on routes and features
- **Lazy Loading**: For non-critical components and libraries
- **Component Memoization**: For expensive renderings

### 7.2. State Update Optimizations
- **Selective Updates**: Update only changed state portions
- **Batch Updates**: Group related state changes
- **Throttled Subscriptions**: Limit real-time update frequency

## 8. Error Handling Patterns

### 8.1. UI Error Boundaries
- **Component-Level Recovery**: Isolate errors to affected components
- **Fallback UI**: Graceful degradation when components fail
- **Error Reporting**: Capture and log client-side errors

### 8.2. API Error Handling
- **Type-Safe Errors**: Structured error responses
- **Retry Logic**: For transient errors
- **User Feedback**: Clear error messages and recovery options

---

This System Patterns document provides a comprehensive reference for the frontend architecture, design patterns, and best practices for the PowerPoint Translator App. It ensures consistency in implementation and serves as a guide for current and future development.
