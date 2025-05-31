# System Patterns: PowerPoint Translator App

## 1. Overview of Architecture
The PowerPoint Translator App follows a **microservice-oriented architecture** for business logic, combined with a **serverless-first** approach for authentication and data persistence via Supabase. The high-level components are:

```
┌───────────────────────────────────────────────────────────────┐
│                       Next.js Frontend                       │
│  • App Router (app/ directory)                                │
│  • Client Components (slide rendering, editing, comments)     │
│  • Server Components for data fetching (using Supabase URL)   │
│  • Zustand store for global state                             │
└───────────────┬───────────────────────────────────────────────┘
                │
                ▼
┌───────────────────────────────────────────────────────────────┐
│                     Supabase (Auth + Postgres)                │
│  • Auth: Email/password, JWT issuance, user management       │
│  • Database: sessions, session_shares, comments, audit_logs   │
│  • Storage: raw PPTX files, optional exported PPTX            │
│  • Realtime: LISTEN/NOTIFY on comments and audit logs         │
│  • RLS: Row-Level Security policies to enforce permissions    │
└───────────────┬───────────────────────────────────────────────┘
                │
                ▼
┌───────────────────────────────────────────────────────────────┐
│                Business Logic Microservices                  │
│  1. SessionService   (Node.js + Supabase SDK)                │
│  2. ChunkService     (Node.js + Supabase SDK)                │
│  3. MergeService     (Node.js + Supabase SDK)                │
│  4. OrderService     (Go + pgx/HTTP to Supabase)             │
│  5. CommentService   (Go + pgx)                               │
│  6. AuditService     (Go + pgx)                               │
│  7. ExportService    (Python + python-pptx + Supabase REST)  │
└───────────────────────────────────────────────────────────────┘
```

## 2. Key Technical Decisions

### 2.1. Supabase for Auth & Data Persistence
- **Auth**: Email/password login using Supabase Auth to issue JWTs.  
- **Postgres**: All core tables—`sessions`, `session_shares`, `comments`, `audit_logs`—live in Supabase.  
- **RLS Policies**: Enforce per-row access so that only session owners or valid share-token holders can read/update data.  
- **Realtime**: Leverage LISTEN/NOTIFY for live comment and audit log updates in the frontend.  

### 2.2. Client-Side PPTX Handling with pptxgenjs
- All parsing of raw PPTX (`.pptx`) to extract slide metadata is done in-browser using **pptxgenjs**.  
- Rendering of slides in the Slide Editor leverages pptxgenjs to build a canvas or SVG representation of each slide.  
- Client-side export: Collect updated metadata and re-create a fully translated PPTX via pptxgenjs in-browser, triggering a download.

### 2.3. Microservice Decomposition
- **SessionService / ChunkService / MergeService**: Use Node.js + TypeScript + `@supabase/supabase-js` to handle JSON manipulation (metadata) and database writes.  
- **OrderService / CommentService / AuditService**: Use Go (with `pgx` or Supabase REST) for structured CRUD operations and strong typing.  
- **ExportService**: Use Python + `python-pptx` for server-side PPTX generation if needed (optional).

### 2.4. Frontend State Management
- **Zustand**: Lightweight store for:
  - Current session ID and owner/reviewer role  
  - Slides metadata and per-slide state (edit buffers, merge selections, reorder state)  
  - Comments cache per shape/run  
  - Notification counts (new comments)  
- **React Context** only for cross-cutting providers (e.g., theme, Supabase client) if required; avoid Redux for speed.

### 2.5. Next.js App Router Patterns
- **Layout File Hierarchy**:
  - `app/layout.tsx`: Root layout with `<SessionProvider>` (Zustand context), navigation UI.  
  - `app/sessions/layout.tsx`: Session-level layout (sidebar + header).  
  - `app/sessions/[sessionId]/layout.tsx`: Session shell with tabs for Slides, History, Share.  
- **Server Components**: For data fetching pages that require SSR (e.g., session list on Dashboard).  
- **Client Components**: For interactive parts (SlideCanvas, TextChunkEditor, CommentThread). Use `"use client"` directive at top.

---

## 3. Design Patterns in Use

### 3.1. Service Layer Pattern
- Abstract all Supabase interactions behind **service modules** (`/lib/services/SessionService.ts`, etc.).  
- Each service exposes methods returning domain models (e.g., `Session`, `SlideMetadata`, `Comment`).  
- Frontend and microservices can share TypeScript interfaces to ensure type safety.

### 3.2. Repository Pattern (Database Access)
- In **Go microservices**, use a `Repository` layer (e.g. `CommentRepository`, `AuditRepository`) that wraps SQL queries (via `pgx`) in well-defined methods (`GetCommentsBySession`, `InsertAuditLog`).  
- This isolates raw SQL and allows easy refactoring if the schema changes.

### 3.3. Adapter/Facade Pattern
- In **Node.js microservices**, wrap Supabase client (`@supabase/supabase-js`) in a **Facade** (`supabaseClient.ts`) that hides configuration and authentication details.  
- Provides simple methods like `supabase.from('sessions').select('metadata')` or `uploadFileToStorage(bucket, path, file)`.

### 3.4. Event Sourcing (Audit Logs)
- Every state-changing action (edit, merge, reorder, comment) writes an immutable entry to `audit_logs`.  
- The app can replay or display history based on these events—ensuring a transparent, chronological record of changes.

### 3.5. Command Pattern (Microservice Endpoints)
- Each microservice endpoint corresponds to a **command**:
  - `UpdateChunkCommand` → triggers `ChunkService.update(...)`.  
  - `MergeCommand` → triggers merging logic.  
  - `ReorderCommand` → triggers order update logic.  
- Commands validate input, perform side effects (DB updates + audit log), and return results.

### 3.6. Observer Pattern (Realtime Notifications)
- **Supabase Realtime** serves as the observer mechanism:  
  - Frontend components subscribe to `comments` and `audit_logs` channels.  
  - On INSERT, Supabase emits events that clients receive in real time.  
  - Zustand store updates observers, and UI badges/timers refresh.

---

## 4. Component Relationships & Interactions

```text
Next.js Frontend
 ├─ pages:
 │    /sessions                  ← Server Component (fetch list of sessions)
 │    /sessions/new              ← Client Component (UploadWizard)
 │    /sessions/[id]/share       ← Client Component (ShareLinkModal)
 │    /sessions/[id]/history     ← Client Component (AuditLogViewer)
 │    /sessions/[id]/slides/[i]  ← Client Component (SlideEditorPage)
 │
 ├─ components:
 │    ├─ UploadWizard
 │    ├─ SlideCanvas             ← uses pptxgenjs to render slide via metadata
 │    ├─ TextChunkEditor         ← overlay on SlideCanvas
 │    ├─ MergePanel              ← uses Zustand to track selected chunks + calls MergeService
 │    ├─ ReadingOrderOverlay     ← drag-and-drop order badges; calls OrderService
 │    ├─ CommentThread           ← fetches via CommentService and subscribes to Realtime
 │    ├─ Toolbar                 ← SaveAll, Merge, ToggleOrder, CommentCenter, Export
 │    └─ NotificationBadge        ← subscribes to Realtime for comment/audit updates
 │
 ├─ hooks:
 │    ├─ useSessionStore         ← Zustand slice: current session, role, etc.
 │    ├─ useSlideStore           ← Zustand slice: slide metadata, updates
 │    ├─ useCommentStore         ← Zustand slice: comments per session/shape
 │    ├─ useNotificationStore    ← Zustand slice: unread counts
 │    ├─ useFetchComments        ← custom hook to call CommentService.getComments
 │    └─ useRealtimeSubscription ← sets up Supabase Realtime listeners
 │
 ├─ lib/services:
 │    ├─ SessionService.ts       ← createSession, finishSession, generateShareToken, validateShareToken
 │    ├─ ChunkService.ts         ← updateChunk(s)
 │    ├─ MergeService.ts         ← mergeRuns
 │    ├─ OrderService.ts         ← reorderSlides
 │    ├─ CommentService.ts       ← getComments, addComment
 │    ├─ AuditService.ts         ← getHistory
 │    └─ ExportService.ts        ← (optional) server-side export orchestration
 │
 ├─ state/
 │    ├─ sessionStore.ts         ← Zustand store for session-level data
 │    ├─ slideStore.ts           ← Zustand store for per-slide state
 │    ├─ commentStore.ts         ← Zustand store for in-session comments
 │    └─ notificationStore.ts    ← Zustand store for real-time notification badges
 │
 └─ utils/
      ├─ supabaseClient.ts       ← config for @supabase/supabase-js (browser + server)
      └─ pptxHelpers.ts          ← helper functions to manipulate pptxgenjs metadata

Backend Microservices
 ├─ SessionService (Node.js)
 │    └─ Interfaces with Supabase Storage & Sessions table
 │
 ├─ ChunkService (Node.js)
 │    └─ Reads/Writes sessions.metadata (JSONB) and audit_logs
 │
 ├─ MergeService (Node.js)
 │    └─ Reconciles runs in metadata, writes audit_logs
 │
 ├─ OrderService (Go)
 │    └─ Updates JSONB order array, writes audit_logs
 │
 ├─ CommentService (Go)
 │    └─ CRUD on comments table, writes audit_logs
 │
 ├─ AuditService (Go)
 │    └─ SELECT on audit_logs table
 │
 └─ ExportService (Python) [Optional]
      └─ Rebuilds PPTX via python-pptx, uploads to Supabase Storage

```

---

## 5. Summary of Patterns & Relationships
1. **Microservices** handle business logic and update Supabase DB + audit.  
2. **Supabase** provides Auth, persistent storage (Postgres + Storage), and Realtime.  
3. **Next.js Frontend** composes data via services, reacts to Realtime events, and uses pptxgenjs for rendering & exporting.  
4. **Zustand** centralizes app state for sessions, slides, comments, and notifications.  
5. **Design Patterns**:  
   - Service Layer & Facade for DB/Storage access  
   - Repository (Go services) for structured queries  
   - Observer (Realtime) for live updates  
   - Command Pattern for microservice endpoints  
   - Event Sourcing via Audit Logs  

This **System Patterns** document serves as the reference for architecture decisions, component interactions, and design patterns in the PowerPoint Translator App. It will help maintain consistency across the codebase and guide future enhancements.  
