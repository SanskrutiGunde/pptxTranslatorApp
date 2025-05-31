<!-- activeContext.md -->

# Active Context: PowerPoint Translator App

## Current Work Focus
As of project initialization, the focus is on establishing the core infrastructure for a web-based PowerPoint translation tool. The system will enable collaborative translation of PPTX presentations while preserving original styling and layout.

## Recent Changes
- **Project Initialization**: Memory bank structure established with core documentation files
- **Architecture Definition**: Microservice-oriented design with Supabase for auth/data and separate services for business logic
- **API Design**: AuditAPI.yaml created for the AuditService microservice specification

## Next Steps
1. **Frontend Setup**
   - Initialize Next.js 14 project with App Router
   - Configure Tailwind CSS and shadcn/ui components
   - Set up Zustand for state management
   - Integrate pptxgenjs for client-side PPTX handling

2. **Supabase Configuration**
   - Set up local Supabase development environment
   - Create database schema (sessions, session_shares, comments, audit_logs)
   - Configure Row-Level Security policies
   - Set up storage buckets for PPTX files

3. **Core Features Implementation**
   - Session creation and PPTX upload flow
   - Authentication with email/password via Supabase Auth
   - Share link generation for reviewer access
   - Basic slide rendering using pptxgenjs

4. **Microservices Development**
   - SessionService (Node.js) for session management
   - ChunkService (Node.js) for text chunk updates
   - CommentService (Go) for collaborative comments
   - AuditService (Go) following the defined API spec

## Active Decisions & Considerations

### Technology Choices
- **Frontend**: Next.js 14 with App Router chosen for its modern React features and SSR capabilities
- **State Management**: Zustand selected over Redux for simplicity and minimal boilerplate
- **PPTX Handling**: Client-side processing with pptxgenjs to avoid server-side parsing complexity
- **Real-time Updates**: Leveraging Supabase Realtime for comments and audit logs

### Architecture Decisions
- **Microservices Split**: Business logic separated by domain (sessions, chunks, comments, etc.)
- **Language Choice**: Node.js for JSON-heavy operations, Go for structured CRUD operations
- **Authentication**: Token-based share links for reviewers to avoid account creation friction

### Development Approach
- **Rapid Prototyping**: Using v0.dev and Cursor AI for scaffolding
- **Internal Tool Focus**: No strict accessibility requirements, desktop-first design
- **Single Developer**: No CI/CD initially, focus on local development

## Current Blockers
None at project initialization phase.

## Key Questions for Next Session
1. Preferred UI component library beyond shadcn/ui?
2. Specific translation languages to support initially?
3. Maximum PPTX file size constraints?
4. Preference for real-time collaboration features beyond comments?

---

*Last Updated: Project Initialization* 