<!-- projectbrief.md -->

# Project Brief: PowerPoint Translator App

## 1. Introduction
The **PowerPoint Translator App** is a web-based internal tool that enables owners and reviewers to collaboratively translate, review, and export PowerPoint presentations. Leveraging Next.js (app directory) for the frontend and Supabase for authentication & data storage, the core value proposition is a seamless, in-browser translation workflow that preserves original styling and layout via pptxgenjs.

## 2. Core Requirements
1. **Session Creation & Management**  
   - Owner can upload a `.pptx` file to create a “translation session.”  
   - System stores the raw PPTX in Supabase Storage and extracts slide metadata (shapes, text runs, positions) into JSON for editing.

2. **Authentication & Access Control**  
   - Owner signs up/logs in via Supabase Auth (email/password).  
   - Owner can generate a secure “share link” (token-based) to invite Reviewers.  
   - Reviewers access via that link (no full account required) and gain read/comment/edit privileges as configured by token.  

3. **Slide Rendering & Editing**  
   - Use pptxgenjs in-browser to render each slide’s shapes and text.  
   - Reviewer can click on any text run to edit the translation.  
   - Reviewer can merge multiple text runs (chunks) into one for better flow.  
   - Reviewer can reorder reading sequence (drag-and-drop) to preserve correct text order for export.

4. **Commenting & Collaboration**  
   - Inline, threaded comments tethered to specific shapes or runs.  
   - Real-time notifications via Supabase Realtime when new comments arrive.  
   - Audit trail logs all edits, merges, reorder actions, and comments.

5. **Export**  
   - Owner can export a final translated PPTX using pptxgenjs (or server-side ExportService) that faithfully preserves slide layout, fonts, and styling.  
   - Downloadable `.pptx` file delivered via browser or Supabase Storage signed URL.

## 3. Goals & Constraints
- **Rapid Development (v0.1)**: Minimize boilerplate; use v0.dev and Cursor AI for scaffolding. Single developer focus; no CI/CD initially.  
- **Internal Use**: No strict accessibility requirements. Support only modern desktop browsers (laptop-first).  
- **In-App Parsing**: Leverage pptxgenjs in the browser to parse metadata—no external parsing service.  
- **User Roles**:  
  - **Owner**: Full control—upload, share, export, view history.  
  - **Reviewer**: Edit translations, merge chunks, reorder text, comment.  
- **Security**: Rely on Supabase RLS for data protection (sessions, comments, audit).  
- **Future-Proofing**: Plan for role-based link permissions, potential server-side PPTX export, and larger concurrency if needed.

## 4. Stakeholders
- **Primary Developer**: Single frontend engineer responsible for initial v0.1 build.  
- **Internal Reviewers**: Subject-matter experts who will translate and comment on slides.  
- **Product Owner**: Collector of requirements, review of scope, and final decision-making on feature priorities.

---
