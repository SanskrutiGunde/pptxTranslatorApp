# Frontend Project Brief: PowerPoint Translator App

## 1. Introduction
The frontend of the **PowerPoint Translator App** is built with Next.js 14 (App Router), focusing on creating an intuitive, responsive UI for translating PowerPoint presentations in-browser. The frontend handles client-side PPTX parsing, slide rendering, collaborative editing, and export while preserving original styling.

## 2. Core Frontend Requirements

1. **Session Management UI**
   - Dashboard to view, create, and access translation sessions
   - PPTX upload interface with drag-and-drop support
   - Session status indicators and progress tracking

2. **Authentication UI**
   - Owner login/signup forms (email/password)
   - Share link generation UI with configurable permissions
   - Reviewer access via token-based links without account creation

3. **Slide Editor Interface**
   - Slide canvas component using pptxgenjs for rendering
   - Text editing overlays for translating individual text runs
   - Intuitive UI for merging text chunks (selection interface)
   - Drag-and-drop reordering of reading sequence

4. **Collaboration Components**
   - Comment threads tethered to specific shapes/runs
   - Real-time notification badges for new comments
   - Commenting toolbar and thread component
   - User presence indicators (optional)

5. **Export Interface**
   - Export settings panel
   - Progress indicator during PPTX generation
   - Download button for translated PPTX

## 3. Frontend Architecture

1. **Next.js App Router**
   - Leveraging file-based routing in app/ directory
   - Server Components for data-fetching pages
   - Client Components for interactive elements

2. **State Management**
   - Zustand stores organized by domain:
     - `sessionStore`: Session metadata, role, permissions
     - `slideStore`: Slide content, edit buffers, selections
     - `commentStore`: Comments cache and thread state
     - `notificationStore`: Real-time notification tracking

3. **UI Component Hierarchy**
   - Layout components for app shell
   - Page-specific components
   - Reusable UI primitives from shadcn/ui

4. **Client-Side PPTX Handling**
   - pptxgenjs for parsing uploaded PPTX files
   - Client-side slide rendering
   - Text extraction and manipulation
   - Export functionality

## 4. Frontend Development Priorities

1. **User Experience**
   - Intuitive slide editor with minimal learning curve
   - Responsive design optimized for desktop/laptop
   - Visual cues for editing, commenting, and collaborating
   - Fast feedback on user actions

2. **Performance Optimization**
   - Efficient slide rendering (canvas/SVG)
   - Pagination for large presentations
   - Careful state management to prevent re-renders
   - Code splitting for optimal bundle size

3. **Developer Experience**
   - Type safety with TypeScript
   - Component organization for maintainability
   - Clear state management patterns
   - Reusable hooks and utilities

## 5. Technical Constraints & Decisions

1. **Browser Support**
   - Modern desktop browsers only (Chrome, Firefox, Safari, Edge)
   - No IE11 or mobile-specific optimizations required

2. **Styling Approach**
   - Tailwind CSS for utility-first styling
   - shadcn/ui components for consistent UI elements
   - No custom CSS files; all styling via Tailwind classes

3. **Client-Side Processing**
   - PPTX parsing happens in-browser with pptxgenjs
   - Slide rendering uses client-side components
   - Export generation can be client-side or server-side

4. **Realtime Features**
   - Comments and notifications use Supabase Realtime
   - Editing is not real-time synchronized (last-write-wins)

---

This Frontend Project Brief outlines the key requirements, architecture decisions, and priorities for the PowerPoint Translator App's frontend implementation. It serves as a guide for frontend development and should be referenced throughout the project lifecycle.
