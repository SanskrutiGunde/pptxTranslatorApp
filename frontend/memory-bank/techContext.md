# Frontend Technology Context: PowerPoint Translator App

## 1. Frontend Technology Stack

### 1.1. Core Technologies
- **Next.js 14**
  - App Router architecture (`app/` directory structure)
  - Server Components for data-fetching
  - Client Components for interactive UI
  - API Routes for backend-less endpoints
  - Built-in image optimization

- **React 18**
  - Functional components with hooks
  - Concurrent rendering features
  - Suspense for data fetching
  - Strict Mode enabled

- **TypeScript 5.1+**
  - Strict type checking enabled
  - Interface-based component props
  - Type-safe state management
  - Comprehensive type definitions

### 1.2. UI & Styling
- **Tailwind CSS**
  - Utility-first CSS framework
  - JIT (Just-In-Time) compiler
  - Custom theme configuration
  - Component-specific design tokens

- **shadcn/ui**
  - Unstyled, accessible UI components
  - Built on Radix UI primitives
  - Customized via Tailwind classes
  - Installed individually as needed

- **UI Animation**
  - Framer Motion for complex animations
  - CSS transitions for simple animations
  - Tailwind's transition utilities

### 1.3. State Management
- **Zustand**
  - Lightweight global state management
  - Slice pattern for domain separation
  - Middleware for persistence (optional)
  - TypeScript integration

- **React Context** (limited use)
  - ThemeProvider
  - SupabaseProvider
  - Other cross-cutting concerns

### 1.4. PPTX Processing
- **pptxgenjs**
  - Client-side PPTX parsing
  - Slide rendering to canvas/SVG
  - Text extraction and manipulation
  - PPTX generation for export

### 1.5. Other Libraries
- **@dnd-kit/core**
  - Drag-and-drop for reading order
  - Accessible interaction patterns

- **date-fns**
  - Date formatting and manipulation
  - Timezone handling

- **zod**
  - Runtime validation for API inputs/outputs
  - Form validation schemas

## 2. Frontend Development Environment

### 2.1. Project Setup
- **Next.js Project Structure**
  ```
  /frontend
    /app
    /components
    /hooks
    /lib
    /styles
    /types
    package.json
    tsconfig.json
    next.config.js
    tailwind.config.js
    postcss.config.js
  ```

- **TypeScript Configuration**
  ```json
  // tsconfig.json
  {
    "compilerOptions": {
      "target": "es2017",
      "lib": ["dom", "dom.iterable", "esnext"],
      "allowJs": true,
      "skipLibCheck": true,
      "strict": true,
      "forceConsistentCasingInFileNames": true,
      "noEmit": true,
      "esModuleInterop": true,
      "module": "esnext",
      "moduleResolution": "node",
      "resolveJsonModule": true,
      "isolatedModules": true,
      "jsx": "preserve",
      "incremental": true,
      "plugins": [
        {
          "name": "next"
        }
      ],
      "paths": {
        "@/*": ["./src/*"]
      }
    },
    "include": ["next-env.d.ts", "**/*.ts", "**/*.tsx", ".next/types/**/*.ts"],
    "exclude": ["node_modules"]
  }
  ```

- **Tailwind Configuration**
  ```javascript
  // tailwind.config.js
  /** @type {import('tailwindcss').Config} */
  module.exports = {
    darkMode: ["class"],
    content: [
      './pages/**/*.{ts,tsx}',
      './components/**/*.{ts,tsx}',
      './app/**/*.{ts,tsx}',
      './src/**/*.{ts,tsx}',
    ],
    theme: {
      container: {
        center: true,
        padding: "2rem",
        screens: {
          "2xl": "1400px",
        },
      },
      extend: {
        colors: {
          border: "hsl(var(--border))",
          input: "hsl(var(--input))",
          ring: "hsl(var(--ring))",
          background: "hsl(var(--background))",
          foreground: "hsl(var(--foreground))",
          primary: {
            DEFAULT: "hsl(var(--primary))",
            foreground: "hsl(var(--primary-foreground))",
          },
          secondary: {
            DEFAULT: "hsl(var(--secondary))",
            foreground: "hsl(var(--secondary-foreground))",
          },
          // ... other color tokens
        },
        borderRadius: {
          lg: "var(--radius)",
          md: "calc(var(--radius) - 2px)",
          sm: "calc(var(--radius) - 4px)",
        },
        keyframes: {
          // ... animation keyframes
        },
        animation: {
          // ... animation definitions
        },
      },
    },
    plugins: [require("tailwindcss-animate")],
  }
  ```

### 2.2. Environment Configuration
- **Environment Variables**
  ```
  # .env.local
  NEXT_PUBLIC_SUPABASE_URL=http://localhost:54321
  NEXT_PUBLIC_SUPABASE_ANON_KEY=<anon-key>
  NEXT_PUBLIC_APP_URL=http://localhost:3000
  
  # Service endpoints
  NEXT_PUBLIC_SESSION_SERVICE_URL=http://localhost:4001
  NEXT_PUBLIC_CHUNK_SERVICE_URL=http://localhost:4002
  NEXT_PUBLIC_MERGE_SERVICE_URL=http://localhost:4003
  NEXT_PUBLIC_ORDER_SERVICE_URL=http://localhost:4004
  NEXT_PUBLIC_COMMENT_SERVICE_URL=http://localhost:4005
  NEXT_PUBLIC_AUDIT_SERVICE_URL=http://localhost:4006
  NEXT_PUBLIC_EXPORT_SERVICE_URL=http://localhost:4007
  ```

- **Next.js Config**
  ```javascript
  // next.config.js
  /** @type {import('next').NextConfig} */
  const nextConfig = {
    reactStrictMode: true,
    images: {
      domains: ['localhost'],
    },
    experimental: {
      serverActions: true,
    },
  };
  
  module.exports = nextConfig;
  ```

### 2.3. Package Dependencies
```json
// Key dependencies in package.json
{
  "dependencies": {
    "next": "^14.0.0",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "@supabase/supabase-js": "^2.21.0",
    "@supabase/auth-helpers-nextjs": "^0.7.0",
    "pptxgenjs": "^3.12.0",
    "zustand": "^4.3.8",
    "tailwindcss": "^3.3.2",
    "@dnd-kit/core": "^6.0.8",
    "@dnd-kit/sortable": "^7.0.2",
    "date-fns": "^2.30.0",
    "zod": "^3.21.4",
    "framer-motion": "^10.12.16",
    "clsx": "^1.2.1",
    "tailwind-merge": "^1.13.2"
  },
  "devDependencies": {
    "typescript": "^5.1.3",
    "@types/react": "^18.2.12",
    "@types/node": "^20.3.1",
    "postcss": "^8.4.24",
    "autoprefixer": "^10.4.14",
    "eslint": "^8.42.0",
    "eslint-config-next": "^13.4.5",
    "prettier": "^2.8.8",
    "prettier-plugin-tailwindcss": "^0.3.0"
  }
}
```

## 3. Development Workflow

### 3.1. Development Commands
```bash
# Start development server
npm run dev

# Type checking
npm run type-check

# Linting
npm run lint

# Build for production
npm run build

# Run production build locally
npm run start

# Format code
npm run format
```

### 3.2. shadcn/ui Component Installation
```bash
# Initialize shadcn/ui
npx shadcn-ui@latest init

# Add individual components
npx shadcn-ui@latest add button
npx shadcn-ui@latest add card
npx shadcn-ui@latest add dialog
# ... etc.
```

### 3.3. Supabase Integration

- **Client Setup**
```typescript
// lib/utils/supabaseClient.ts
import { createClientComponentClient, createServerComponentClient } from '@supabase/auth-helpers-nextjs';
import { cookies } from 'next/headers';
import { Database } from '@/types/supabase';

// For client components
export const createClient = () => {
  return createClientComponentClient<Database>({
    supabaseUrl: process.env.NEXT_PUBLIC_SUPABASE_URL!,
    supabaseKey: process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!,
  });
};

// For server components
export const createServerClient = () => {
  return createServerComponentClient<Database>({
    cookies,
  });
};
```

- **Auth Integration**
```typescript
// app/auth/login/page.tsx
'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { createClient } from '@/lib/utils/supabaseClient';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();
  const supabase = createClient();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    
    const { error } = await supabase.auth.signInWithPassword({
      email,
      password,
    });
    
    if (error) {
      setError(error.message);
    } else {
      router.push('/sessions');
      router.refresh();
    }
  };
  
  return (
    // Login form JSX
  );
}
```

- **Realtime Subscriptions**
```typescript
// hooks/useRealtimeSubscription.ts
import { useEffect } from 'react';
import { createClient } from '@/lib/utils/supabaseClient';
import { useCommentStore } from '@/hooks/useCommentStore';
import { useNotificationStore } from '@/hooks/useNotificationStore';

export function useRealtimeSubscription(sessionId: string) {
  const addComment = useCommentStore((state) => state.addComment);
  const incrementNotification = useNotificationStore((state) => state.increment);
  const supabase = createClient();
  
  useEffect(() => {
    const channel = supabase
      .channel(`comments:${sessionId}`)
      .on('postgres_changes', {
        event: 'INSERT',
        schema: 'public',
        table: 'comments',
        filter: `session_id=eq.${sessionId}`,
      }, (payload) => {
        addComment(payload.new);
        incrementNotification('comments');
      })
      .subscribe();
      
    return () => {
      supabase.removeChannel(channel);
    };
  }, [sessionId, addComment, incrementNotification, supabase]);
}
```

### 3.4. PPTX Processing Implementation

- **File Upload & Parsing**
```typescript
// hooks/usePPTXUpload.ts
export function usePPTXUpload() {
  const [uploading, setUploading] = useState(false);
  const [progress, setProgress] = useState(0);
  const sessionService = new SessionService();
  
  const uploadFile = async (file: File, sessionName: string) => {
    setUploading(true);
    setProgress(0);
    
    try {
      // 1. Parse the PPTX with pptxgenjs to extract metadata
      const pptxgen = await import('pptxgenjs');
      const pptx = new pptxgen.Presentation();
      
      // Load the file and extract slide metadata
      // This is a simplified representation
      await pptx.load(file);
      
      const slideMetadata = [];
      for (let i = 0; i < pptx.slides.length; i++) {
        const slide = pptx.slides[i];
        // Extract text elements, shapes, etc.
        slideMetadata.push({
          slideNumber: i + 1,
          shapes: [], // Extracted shapes and text runs
        });
        setProgress((i + 1) / pptx.slides.length * 50);
      }
      
      // 2. Create session and upload file
      const session = await sessionService.createSession(
        sessionName,
        file,
        slideMetadata
      );
      
      setProgress(100);
      return session;
    } catch (error) {
      console.error('PPTX upload error:', error);
      throw error;
    } finally {
      setUploading(false);
    }
  };
  
  return { uploadFile, uploading, progress };
}
```

## 4. Frontend Technical Constraints

### 4.1. Browser Support
- **Target Browsers**
  - Chrome 90+
  - Firefox 90+
  - Safari 14+
  - Edge 90+
- **No IE11 Support Required**
- **No Mobile-Specific Optimizations Required**

### 4.2. Performance Targets
- **First Contentful Paint**: < 1.2s
- **Time to Interactive**: < 2.5s
- **Slide Rendering Time**: < 500ms per slide
- **Memory Usage**: < 200MB for typical presentations
- **Bundle Size**: < 500KB (initial load, gzipped)

### 4.3. Technical Limitations
- **PPTX Size**: Optimal for files < 10MB
- **Slide Complexity**: Some complex PowerPoint features may have simplified rendering
- **Font Compatibility**: Limited to web-safe fonts or embedded fonts
- **Animation Support**: Static rendering only (no PowerPoint animations)

## 5. Frontend Development Best Practices

### 5.1. Code Organization
- **File Naming**
  - PascalCase for React components (`SlideCanvas.tsx`)
  - camelCase for utilities, hooks, services (`pptxHelpers.ts`, `useSessionStore.ts`)
  - One component per file

- **Import Order**
  1. React/Next.js imports
  2. Third-party libraries
  3. Project imports (absolute paths)
  4. Local imports (relative paths)
  5. Type imports
  6. CSS imports

### 5.2. Component Patterns
- **Props Interface**
  ```typescript
  interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    variant?: 'default' | 'outline' | 'ghost';
    size?: 'sm' | 'md' | 'lg';
    isLoading?: boolean;
  }
  ```

- **Component Declaration**
  ```typescript
  export function Button({
    variant = 'default',
    size = 'md',
    isLoading = false,
    children,
    className,
    ...props
  }: ButtonProps) {
    return (
      <button
        className={cn(
          buttonVariants({ variant, size }),
          isLoading && 'opacity-70 pointer-events-none',
          className
        )}
        disabled={isLoading}
        {...props}
      >
        {isLoading && <Spinner className="mr-2 h-4 w-4" />}
        {children}
      </button>
    );
  }
  ```

### 5.3. Error Handling
- **Form Validation Errors**
  ```typescript
  // Using zod
  const formSchema = z.object({
    email: z.string().email('Invalid email address'),
    password: z.string().min(8, 'Password must be at least 8 characters'),
  });
  ```

- **API Error Handling**
  ```typescript
  // Consistent error response type
  interface ApiError {
    code: string;
    message: string;
    details?: Record<string, any>;
  }
  
  // Error handling
  try {
    const result = await apiCall();
    // Handle success
  } catch (error) {
    if (error instanceof ApiError) {
      // Handle structured API error
    } else {
      // Handle unexpected error
    }
  }
  ```

### 5.4. Styling Conventions
- **Tailwind Class Organization**
  1. Layout (position, display, width, height)
  2. Spacing (margin, padding)
  3. Typography (font, text)
  4. Visual (colors, borders, shadows)
  5. Interactive (hover, focus)
  6. Responsive modifiers

- **Custom Utility Classes**
  ```css
  @layer utilities {
    .text-balance {
      text-wrap: balance;
    }
    
    .scrollbar-hide {
      scrollbar-width: none;
      &::-webkit-scrollbar {
        display: none;
      }
    }
  }
  ```

## 6. Testing Approach

### 6.1. Unit Testing
- **Component Testing**
  - React Testing Library
  - Jest for assertions
  - Mock service layer

- **Hook Testing**
  - `renderHook` from React Testing Library
  - State assertions

### 6.2. Integration Testing
- **Page Testing**
  - Mock Supabase responses
  - Test user flows

### 6.3. E2E Testing (Optional)
- **Playwright**
  - Cross-browser testing
  - Test critical user journeys

## 7. Frontend Performance Considerations

### 7.1. Bundle Size Optimization
- **Dynamic Imports**
  ```typescript
  const PPTXViewer = dynamic(() => import('@/components/PPTXViewer'), {
    loading: () => <LoadingSpinner />,
    ssr: false,
  });
  ```

- **Tree Shaking**
  - ESM imports for better tree shaking
  - Avoid importing entire libraries

### 7.2. Render Performance
- **Memoization**
  ```typescript
  const MemoizedComponent = memo(MyComponent, (prevProps, nextProps) => {
    return prevProps.id === nextProps.id;
  });
  ```

- **Virtualization**
  ```typescript
  import { useVirtualizer } from '@tanstack/react-virtual';
  
  function VirtualizedList({ items }) {
    const parentRef = useRef(null);
    
    const rowVirtualizer = useVirtualizer({
      count: items.length,
      getScrollElement: () => parentRef.current,
      estimateSize: () => 50,
    });
    
    return (
      <div ref={parentRef} className="h-96 overflow-auto">
        <div
          style={{
            height: `${rowVirtualizer.getTotalSize()}px`,
            position: 'relative',
          }}
        >
          {rowVirtualizer.getVirtualItems().map((virtualRow) => (
            <div
              key={virtualRow.index}
              style={{
                position: 'absolute',
                top: 0,
                left: 0,
                width: '100%',
                height: `${virtualRow.size}px`,
                transform: `translateY(${virtualRow.start}px)`,
              }}
            >
              {items[virtualRow.index]}
            </div>
          ))}
        </div>
      </div>
    );
  }
  ```

---

This Frontend Technology Context document serves as a comprehensive technical reference for the PowerPoint Translator App's frontend implementation. It outlines the technology stack, development environment, code organization, and best practices to ensure consistent and efficient frontend development.
