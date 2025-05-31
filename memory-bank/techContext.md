# Technology Context: PowerPoint Translator App

## 1. Overview of Technology Stack

### 1.1. Frontend
- **Next.js 14 (App Router)**  
  - React-based framework with file-system–based routing under `app/`.  
  - Supports Server Components and Client Components via `"use client"` directive.  
  - Built-in Fast Refresh, CSS support, and image optimization.

- **React + TypeScript**  
  - Typed React components enhance reliability and catch errors early.  
  - All code lives under `src/` with `tsconfig.json` enabling strict type checking.

- **Zustand**  
  - Lightweight state management library with no provider boilerplate.  
  - Used to store global state slices: session, slide, comment, and notification data.

- **Tailwind CSS**  
  - Utility-first CSS framework with JIT compiler.  
  - Configured in `tailwind.config.js`, all styling via class names.  
  - Global import in `globals.css`.

- **shadcn/ui**  
  - Collection of accessible UI primitives (Buttons, Inputs, Cards, Dialogs, etc.) built on top of Radix UI.  
  - Initialized via `npx shadcn-ui init`; components live in `/components/ui`.

- **pptxgenjs**  
  - Client-side library for reading, rendering, and writing PPTX files.  
  - Used in-browser for slide parsing, rendering (SlideCanvas), and final export.

### 1.2. Backend (Data & Auth)
- **Supabase**  
  - **Auth**: Email/password, JWT issuance.  
  - **Database**: PostgreSQL with Row-Level Security (RLS) for data isolation.  
  - **Storage**: Buckets for raw PPTX (`raw-pptx`) and optional exported PPTX (`exports`).  
  - **Realtime**: LISTEN/NOTIFY on `comments` and `audit_logs` for live updates in the UI.  

### 1.3. Microservices
- **Node.js (TypeScript)**  
  - SessionService, ChunkService, MergeService.  
  - Uses `@supabase/supabase-js` for authenticated DB/storage access.  
  - Runs on Express or Fastify.

- **Go**  
  - OrderService, CommentService, AuditService.  
  - Uses `github.com/jackc/pgx/v5` for Postgres JSONB updates and queries.  
  - Compiled static binaries, minimal runtime dependencies.

- **Python (Optional)**  
  - ExportService to rebuild PPTX on server using `python-pptx`.  
  - Uses Supabase REST endpoints or `supabase-py` to upload results to storage.

---

## 2. Development Setup

### 2.1. Prerequisites
- **Node.js v18+** (LTS)  
- **npm or Yarn** for package management  
- **Go 1.20+** (for Go microservices)  
- **Python 3.11+** (if using ExportService)  
- **Supabase CLI** (`npm install -g supabase`)  
- **Git** for version control  

### 2.2. Supabase Local Dev
1. Install Supabase CLI and Docker.  
2. In project root, run:
   ```bash
   supabase init
   supabase start
   ```
   This starts local Postgres, Auth, and Storage emulators.  
3. Create schema/files:  
   - `supabase/migrations/` for table definitions (create `sessions`, `session_shares`, `comments`, `audit_logs`).  
   - `supabase/policies/` for RLS policy SQL.  
   - `supabase/functions/` for any Supabase Edge Functions.  
4. Push schema to local DB:  
   ```bash
   supabase db push
   ```
5. Generate TypeScript types (optional):  
   ```bash
   supabase gen types typescript --project-id <your-project> > supabase/types.ts
   ```

### 2.3. Next.js Frontend Setup
1. Clone repo, then:
   ```bash
   cd <project-root>
   npm install
   ```
2. Create `.env.local`:
   ```env
   NEXT_PUBLIC_SUPABASE_URL=http://localhost:54321
   NEXT_PUBLIC_SUPABASE_ANON_KEY=<anon-key-from-supabase>
   SUPABASE_SERVICE_ROLE_KEY=<service-role-key-from-supabase>
   SESSION_SERVICE_URL=http://localhost:4001
   CHUNK_SERVICE_URL=http://localhost:4002
   MERGE_SERVICE_URL=http://localhost:4003
   ORDER_SERVICE_URL=http://localhost:4004
   COMMENT_SERVICE_URL=http://localhost:4005
   AUDIT_SERVICE_URL=http://localhost:4006
   EXPORT_SERVICE_URL=http://localhost:4007   # optional
   ```
3. Tailwind CSS initialization (if not included by v0.dev):
   ```bash
   npx tailwindcss init -p
   ```
   - `tailwind.config.js` should include:
     ```js
     module.exports = {
       content: [
         './app/**/*.{js,ts,jsx,tsx}',
         './components/**/*.{js,ts,jsx,tsx}',
       ],
       theme: { extend: {} },
       plugins: [],
     };
     ```
   - Add import to `globals.css`:
     ```css
     @tailwind base;
     @tailwind components;
     @tailwind utilities;
     ```

4. Run Next.js dev server:
   ```bash
   npm run dev
   ```
   - Frontend accessible at `http://localhost:3000`.

### 2.4. Microservices Setup
- **Node.js Services (SessionService, ChunkService, MergeService)**:
  1. Navigate to each service folder (e.g. `services/session-service`).
  2. Install dependencies:
     ```bash
     npm install
     ```
  3. Create a local `.env`:
     ```env
     SUPABASE_URL=http://localhost:54321
     SUPABASE_SERVICE_ROLE_KEY=<service-role-key>
     ```
  4. Start the service:
     ```bash
     npm run dev   # or `ts-node src/index.ts`
     ```
- **Go Services (OrderService, CommentService, AuditService)**:
  1. Navigate to each service folder (e.g. `services/order-service`).
  2. Create a local `.env`:
     ```env
     SUPABASE_URL=http://localhost:54321
     SUPABASE_SERVICE_ROLE_KEY=<service-role-key>
     ```
  3. Run:
     ```bash
     go run main.go
     ```
- **Python Service (ExportService)** (optional):
  1. Navigate to `services/export-service`.
  2. Create virtual environment and install:
     ```bash
     python -m venv venv
     source venv/bin/activate
     pip install fastapi uvicorn python-pptx supabase-py
     ```
  3. Start:
     ```bash
     uvicorn main:app --reload
     ```

---

## 3. Technical Constraints & Dependencies

### 3.1. Constraints
- **Internal‐Only**: No need for public hosting; run everything locally (no CI/CD initially).  
- **File Size**: PPTX files average ~8 MB; ensure in-browser parsing and file uploads handle this comfortably.  
- **No Accessibility Mandate**: Focus on desktop‐first design, no WCAG compliance required.  
- **Real-Time Comments**: Rely on Supabase Realtime; do not build custom WebSocket server.

### 3.2. Dependencies
- **Frontend**:
  - `next`, `react`, `react-dom`  
  - `@supabase/supabase-js`  
  - `zustand`  
  - `tailwindcss`, `postcss`, `autoprefixer`  
  - `shadcn/ui` (via `npx shadcn-ui init`)  
  - `pptxgenjs`  
  - `@dnd-kit/core` (for drag-and-drop reorder)  
  - (Optional) `react-query` or `swr` if caching data beyond Zustand is desired  

- **Node.js Microservices**:
  - `@supabase/supabase-js`  
  - `express` or `fastify`  
  - `typescript`, `ts-node` (if using TS)  
  - `uuid` (for token/session ID generation)  

- **Go Microservices**:
  - `github.com/jackc/pgx/v5` (Postgres driver)  
  - `github.com/go-chi/chi/v5` (router) or `net/http` (standard library)  
  - `github.com/google/uuid` (UUID generation)  

- **Python ExportService**:
  - `fastapi`  
  - `uvicorn`  
  - `python-pptx`  
  - `supabase-py`  

### 3.3. Local Environment Tools
- **Supabase CLI**  
  - Local Postgres + Auth + Storage for development.  
  - Commands: `supabase init`, `supabase start`, `supabase db push`, `supabase gen types`.  

- **Node.js & NPM/Yarn**  
  - For frontend and Node microservices.  

- **Go Toolchain**  
  - For building/running Go microservices.  

- **Python 3.11+ & uv**  
  - For optional ExportService.  

---

## 4. Summary of Tech Context
1. **Next.js Frontend** with React, TypeScript, Tailwind CSS, shadcn/ui, pptxgenjs, and Zustand.  
2. **Supabase** as the sole provider of Auth (JWT), Storage (raw PPTX), Postgres DB (sessions, comments, audit), and Realtime.  
3. **Microservices** split by business concern: JSON‐heavy logic in Node.js/TypeScript; structured CRUD in Go; optional server-side PPTX export in Python.  
4. **Local‐Only Setup**: Supabase CLI for local environment, no CI/CD or Docker required.  
5. **Dependencies**: Lean package sets for each service, emphasizing rapid development and minimal operational overhead.  

This **Tech Context** document is the reference for tools, setup instructions, dependencies, and constraints for the PowerPoint Translator App. It ensures all developers share a clear understanding of the chosen stack and can get started quickly.  
