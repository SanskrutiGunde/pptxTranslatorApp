# Frontend Product Context: PowerPoint Translator App

## 1. Frontend User Experience Goals

### 1.1. Owner Experience
- **Session Creation Flow**
  - Simple, intuitive dashboard for managing translation sessions
  - Drag-and-drop PPTX upload with progress indicator
  - Clear status indicators for in-progress translations
  - Ability to easily generate and copy share links

- **Session Management**
  - At-a-glance view of session status, collaborators, and progress
  - Clear indicators for comments requiring attention
  - Quick access to history and audit logs
  - Streamlined export process with format options

### 1.2. Reviewer Experience
- **First-Time Access**
  - Minimal friction: click link → enter name → start translating
  - No account creation required (token-based access)
  - Optional brief onboarding tooltip tour

- **Translation Workflow**
  - Slide navigator with thumbnail previews
  - Text runs clearly indicated with visual highlight on hover
  - Inline text editor that maintains position context
  - Visual indicators for merged/reordered text

- **Collaboration Features**
  - Comment icon adjacent to each text element
  - Threaded comment interface with notifications
  - Clear visual indication of which elements have comments
  - Real-time comment count updates

## 2. Frontend Interface Components

### 2.1. Dashboard & Session Management
- **Session Dashboard**
  - Card-based session list with thumbnails
  - Status badges (Draft, In Progress, Ready for Export)
  - Quick action buttons (Share, Export, Delete)

- **Upload Wizard**
  - Multi-step flow: Upload → Configure → Share
  - File validation and preview
  - Language selection (source and target)

- **Share Management**
  - Generated links with copy-to-clipboard
  - Permission toggles (Can Edit, Can Comment)
  - Optional password protection
  - Link expiration settings

### 2.2. Slide Editor
- **Slide Navigator**
  - Thumbnail strip with pagination
  - Current slide indicator
  - Progress indicator (translated/total elements)

- **Slide Canvas**
  - Rendered slide at actual proportions
  - Interactive text elements with hover states
  - Support for text boxes, tables, and shapes
  - Maintains original styling and position

- **Text Editor**
  - Context-aware popup editor
  - Side-by-side original/translated text
  - Character count and overflow warning
  - Format preservation controls

- **Merge Interface**
  - Checkbox selection for multiple text runs
  - Merge button with preview
  - Warning for style inconsistencies
  - Undo merge capability

- **Reading Order**
  - Numbered badges overlaid on text elements
  - Drag-and-drop reordering
  - Visual flow indicators
  - Toggle to show/hide order badges

### 2.3. Collaboration Tools
- **Comment System**
  - Threaded comments with reply capability
  - User attribution and timestamps
  - Comment status (Open, Resolved)
  - Comment counter badge in UI

- **Notification Center**
  - Badge counters for new comments
  - Dropdown notification list
  - Quick navigation to commented elements
  - Read/unread status tracking

### 2.4. Export Interface
- **Export Settings**
  - Format options (if applicable)
  - Slide range selection
  - Include/exclude comments option
  - Quality settings (if applicable)

- **Export Progress**
  - Visual progress indicator
  - Cancel option
  - Status messages
  - Download button on completion

## 3. Frontend User Flows

### 3.1. Owner Flow
1. **Login/Signup**
   - Email/password authentication
   - Remember me option
   - Password reset flow

2. **Session Creation**
   - Dashboard → "New Session" button
   - Upload PPTX (drag-drop or file picker)
   - Set session name and language pair
   - Create session (triggers parsing and metadata extraction)

3. **Sharing**
   - Navigate to Share tab
   - Configure permissions
   - Generate and copy link
   - Optionally set password or expiration

4. **Monitoring**
   - Dashboard view of all sessions
   - Click into session for detailed progress
   - View comment counts and activity
   - Check history for audit trail

5. **Export**
   - Navigate to Export tab
   - Configure export settings
   - Generate translated PPTX
   - Download file

### 3.2. Reviewer Flow
1. **Access**
   - Click share link
   - Enter name (and optional password if set)
   - Land directly in slide editor

2. **Translation**
   - Navigate slides via thumbnail strip
   - Click text element to edit translation
   - Type translated text
   - Save changes (auto-save or manual)

3. **Advanced Editing**
   - Select multiple text runs
   - Click "Merge" to combine
   - Toggle reading order mode
   - Drag-drop numbers to reorder

4. **Collaboration**
   - Click comment icon on text element
   - Add comment or reply
   - Receive notifications for replies
   - Mark comments as resolved

## 4. Frontend Technical Considerations

### 4.1. Responsive Design
- **Breakpoints**
  - Desktop-first design (minimum 1280px)
  - Limited tablet support (1024px)
  - No specific mobile optimization required

- **Layout Adaptability**
  - Flexible slide canvas with zoom controls
  - Collapsible sidebar for smaller screens
  - Responsive component sizing with Tailwind

### 4.2. Performance Optimizations
- **Slide Rendering**
  - Virtualized list for thumbnail strip
  - Only render visible slides
  - Canvas/SVG optimization for complex slides

- **State Management**
  - Minimize re-renders with selective store updates
  - Debounced text editing saves
  - Memoization of complex components

- **Asset Loading**
  - Lazy load pptxgenjs library
  - Preload next/previous slides
  - Optimize image assets in slides

### 4.3. Error Handling & Recovery
- **User-Friendly Errors**
  - Contextual error messages
  - Guided recovery steps
  - Non-blocking notifications

- **Auto-Save**
  - Periodic state persistence
  - Draft recovery on session re-entry
  - Local storage fallback for unsaved changes

- **Connection Handling**
  - Offline indicators
  - Reconnection strategies
  - Conflict resolution on reconnect

## 5. Frontend Success Metrics

- **Usability Metrics**
  - Time to first edit after session creation
  - Average time per slide translation
  - Error rates in translation workflow
  - Number of UI-related support requests

- **Performance Metrics**
  - Initial load time for slide editor
  - Time to parse and render PPTX
  - Export generation time
  - Memory usage during editing

- **Collaboration Metrics**
  - Comment response time
  - Number of comment threads per session
  - Usage patterns for merge and reorder features
  - Time savings compared to manual translation methods

---

This Frontend Product Context document serves as a guide for implementing the user interface and experience of the PowerPoint Translator App. It defines the key interactions, components, and flows that should be prioritized during development to create an intuitive and efficient translation tool.
