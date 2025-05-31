<!-- productContext.md -->

# Product Context: PowerPoint Translator App

## 1. Why This Project Exists
Organizations working with global teams often need to translate existing PowerPoint presentations from English into other languages (Spanish, French, etc.). Current workflows typically involve:
- Exporting to Word or text documents, running batch translation, and manually reassembling slides.
- Sending PPTX files to separate translation agencies, which introduces delays, errors, and loss of slide formatting.
- Lack of a unified, in-browser tool for translators to see the original slide design while editing translated text.

**The PowerPoint Translator App** solves these pain points by providing a dedicated, collaborative environment where translators and reviewers can:
- See the original slide layout (shapes, images, text boxes) in real time.
- Edit translated text in situ without jeopardizing layout.
- Merge or reorder text chunks for natural language flow.
- Comment and discuss specific slide elements inline.
- Export a high-fidelity translated PPTX with one click.

## 2. Problems to Solve
1. **Preserve Slide Fidelity**  
   Translators need to maintain original design elements—fonts, positions, text boxes—while editing.  
   - Traditional methods often break layouts or require manual reformatting.

2. **Streamline Collaboration**  
   Teams of translators and reviewers currently use email or shared drives, leading to version confusion, delayed feedback, and disconnected tools.  
   - There is no in-app comment or audit trail tied directly to specific shapes or runs.

3. **Eliminate Context Switching**  
   Translators toggle between PowerPoint and external translation tools or spreadsheets. This interrupts focus and invites errors.  
   - A single “pane of glass” solution reduces cognitive load and mistakes.

4. **Simplify Export**  
   After translation, reassembling slides or manually replacing text is time-consuming.  
   - A direct export from within the app ensures that translated text is embedded correctly, maintaining formatting.

## 3. Key User Experience Goals
1. **Intuitive Slide Editing**  
   - Render actual slide shapes in the browser using pptxgenjs.  
   - Overlay editable text fields exactly where the original text appears.  
   - Allow merging of text chunks via simple checkboxes and a “Merge” button.  
   - Enable drag-and-drop of numbered badges to adjust reading order.

2. **Seamless Commenting Workflow**  
   - Each text run or shape should have a small “comment” icon.  
   - Clicking opens a threaded discussion panel for that specific element.  
   - Notifications (badge indicator) alert users to new comments in real time.

3. **Clear Role Separation: Owner vs. Reviewer**  
   - Owner UI: Focus on uploading, sharing, and exporting a presentation.  
   - Reviewer UI: Focus on translating, merging, reordering, and commenting.  
   - Minimal login friction for reviewers via token-based share link.

4. **Real-Time Feedback & History**  
   - Changes (edits, merges, reorder) reflect immediately in the UI.  
   - Audit log accessible to both roles shows a chronological record of all actions.  
   - History page displays user, action type, timestamp, and details.

5. **Minimal Setup & Rapid Onboarding**  
   - Owner needs only to sign up and drag-drop a PPTX to start.  
   - Reviewer clicks a share link, enters their name or an optional password, and begins editing—no separate account creation required.  
   - Provide tooltips or a short guided walkthrough for first-time users.

## 4. Core Value Proposition
- **Time Savings**: Translate entire slide decks without leaving the browser.  
- **Reduced Errors**: Inline editing prevents layout breakage and ensures consistent formatting.  
- **Enhanced Collaboration**: Threaded comments and shared sessions keep teams in sync.  
- **Single Source of Truth**: All versions and audit logs reside in one platform (Supabase), eliminating file sprawl.

---