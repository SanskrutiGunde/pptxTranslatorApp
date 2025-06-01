# v0.dev Prompt for PowerPoint Translator App Frontend MVP

## Application Overview
Create a modern, intuitive web application for translating PowerPoint presentations. The app allows users to upload PPTX files, translate text content while preserving original formatting, collaborate via comments, and export the translated presentation. Built with Next.js 14, Tailwind CSS, shadcn/ui components, and Supabase for authentication and storage.

## Design System
- Use a clean, professional interface with subtle animations
- Primary color: blue (#3B82F6)
- Secondary color: slate (#64748B)
- Accent colors: emerald green for success (#10B981), amber for warnings (#F59E0B), rose for errors (#F43F5E)
- Font: Inter
- Rounded corners for cards, buttons, and inputs (medium radius)
- Subtle drop shadows for elevated components
- Consistent spacing and padding throughout the application
- Desktop-first design (min-width 1280px), with basic support for tablets

## Core UI Components

### Authentication
Create a polished login/signup page with:
- Email/password login form with validation
- "Remember me" checkbox
- Password reset link
- Clean error handling
- Supabase authentication integration (show with mock functionality)
- Subtle branding for PowerPoint Translator App

### Dashboard
Design a dashboard for managing translation sessions with:
- Header with user profile menu (logout option)
- "New Session" prominent button
- Session cards in a grid layout showing:
  - Session name
  - Creation date
  - Status badge (Draft, In Progress, Ready for Export)
  - Thumbnail preview of first slide
  - Progress indicator (% translated)
  - Quick action buttons (Share, Export, Delete)
- Empty state for no sessions

### Upload Wizard
Design a multi-step upload flow:
1. **Upload Step**:
   - Drag-and-drop zone for PPTX files
   - File browser button alternative
   - Upload progress indicator
   - File validation messaging

2. **Configure Step**:
   - Session name input
   - Language pair selection (source/target dropdown)
   - Parsing progress indicator showing slide extraction

3. **Success Step**:
   - Success confirmation
   - Preview of first slide
   - Options to view slides or share now

### Slide Editor
Create the main slide editing interface with:
1. **Three-column layout**:
   - Left sidebar: Slide navigator with thumbnails
   - Center: Main slide canvas
   - Right sidebar: Comments panel

2. **Slide Navigator**:
   - Vertical list of slide thumbnails
   - Current slide indicator
   - Progress indicator per slide
   - Pagination for large presentations

3. **Slide Canvas**:
   - Rendered PowerPoint slide (show a realistic example)
   - Text elements with hover highlights
   - Interactive text overlays for editing
   - Visual indicators for commented elements

4. **Text Editing Interface**:
   - Inline editor appearing when text element is clicked
   - Original text shown above (read-only)
   - Translation input field below
   - Save/cancel buttons
   - Character count

5. **Merge Interface**:
   - Checkbox selection for multiple text runs
   - "Merge Selected" button
   - Preview of merged content

6. **Reading Order Interface**:
   - Numbered badges on text elements
   - Drag handles for reordering
   - Toggle to show/hide reading order

### Comments & Collaboration
Design the commenting system with:
1. **Comment Thread**:
   - List of comments attached to text element
   - User attribution and timestamps
   - Reply functionality
   - Resolve/reopen toggles

2. **Comment Form**:
   - Text input for new comments
   - Submit button
   - Mention functionality (optional)

3. **Notification System**:
   - Badge counters on commented elements
   - Notification dropdown in header
   - Unread/read status indicators

### Export Interface
Create an export panel with:
- Format options
- Slide range selection
- Export progress indicator
- Download button
- Success/error messaging

## Specific Component Details

### SessionCard
```
Component: SessionCard
Description: Card displaying translation session information on the dashboard
Properties:
- session: {
    id: string
    name: string
    createdAt: Date
    status: 'draft' | 'in-progress' | 'ready'
    progress: number
    thumbnailUrl: string
    slideCount: number
  }
- onShare: function
- onExport: function
- onDelete: function
```

### SlideCanvas
```
Component: SlideCanvas
Description: Main component that renders a PowerPoint slide and overlays interactive elements
Properties:
- slide: {
    id: string
    number: number
    shapes: Array<{
      id: string
      type: string
      x: number
      y: number
      width: number
      height: number
      text?: string
      translatedText?: string
      hasComments: boolean
    }>
    background?: string
  }
- editable: boolean
- onTextClick: function
- showReadingOrder: boolean
```

### TextChunkEditor
```
Component: TextChunkEditor
Description: Popup editor for translating text chunks
Properties:
- originalText: string
- translatedText: string
- position: { x: number, y: number }
- onSave: function
- onCancel: function
- maxLength?: number
```

### UploadWizard
```
Component: UploadWizard
Description: Multi-step wizard for uploading and configuring PPTX files
Properties:
- onComplete: function
- supportedLanguages: string[]
```

### CommentThread
```
Component: CommentThread
Description: Displays threaded comments for a specific element
Properties:
- comments: Array<{
    id: string
    userId: string
    userName: string
    content: string
    createdAt: Date
    isResolved: boolean
    parentId?: string
  }>
- onAddComment: function
- onResolve: function
```

## Interactive Features to Demonstrate

1. **File Upload**: Show drag-and-drop functionality with progress indicator
2. **Text Editing**: Display popup editor when clicking on a text element
3. **Merge Functionality**: Show selection and merging of multiple text runs
4. **Reading Order**: Demonstrate drag-and-drop reordering of numbered elements
5. **Comments**: Show adding and resolving comments
6. **Share Link**: Display generating and copying a share link
7. **Export**: Show export flow with progress and download

## Page Examples to Create

1. **Login Page**: Clean authentication screen
2. **Dashboard**: Session management overview
3. **Upload Wizard**: All three steps of the upload process
4. **Slide Editor**: Full editing interface with all panels
5. **Share Page**: Interface for generating and managing share links
6. **Export Page**: Export configuration and progress interface

## Animation & Interaction Notes

- Smooth transitions between wizard steps
- Subtle hover effects on interactive elements
- Progress indicators with animation for async operations
- Micro-interactions for successful actions (saves, comments)
- Loading skeletons for asynchronous content

## Technical Considerations

- Components should look like they're built with Next.js 14 and Tailwind
- All UI elements should be accessible
- Responsive down to 1024px width (tablet)
- Error states should be handled gracefully
- Empty states should be designed for all lists
- Optimistic UI updates for better user experience

---

## Example Data for Visualization

### Sample Session
```json
{
  "id": "sess_123456",
  "name": "Q2 Marketing Presentation",
  "createdAt": "2025-05-28T14:30:00Z",
  "status": "in-progress",
  "progress": 65,
  "slideCount": 24,
  "language": {
    "source": "English",
    "target": "Spanish"
  }
}
```

### Sample Slide
```json
{
  "id": "slide_789012",
  "number": 5,
  "shapes": [
    {
      "id": "shape_1",
      "type": "title",
      "x": 0.1,
      "y": 0.1,
      "width": 0.8,
      "height": 0.1,
      "text": "Q2 Marketing Results",
      "translatedText": "Resultados de Marketing del Q2",
      "hasComments": true
    },
    {
      "id": "shape_2",
      "type": "subtitle",
      "x": 0.1,
      "y": 0.25,
      "width": 0.8,
      "height": 0.05,
      "text": "Performance across channels",
      "translatedText": "Rendimiento a través de canales",
      "hasComments": false
    },
    {
      "id": "shape_3",
      "type": "bullet",
      "x": 0.1,
      "y": 0.35,
      "width": 0.8,
      "height": 0.05,
      "text": "Social media engagement increased by 24%",
      "translatedText": "El compromiso en redes sociales aumentó un 24%",
      "hasComments": false
    }
  ],
  "background": "#ffffff"
}
```

### Sample Comments
```json
[
  {
    "id": "comment_1",
    "userId": "user_123",
    "userName": "Alex Johnson",
    "content": "Should we use 'resultados trimestrales' instead for a more formal tone?",
    "createdAt": "2025-05-29T09:15:00Z",
    "isResolved": false,
    "elementId": "shape_1"
  },
  {
    "id": "comment_2",
    "userId": "user_456",
    "userName": "Maria Garcia",
    "content": "Good suggestion, I've updated it.",
    "createdAt": "2025-05-29T10:22:00Z",
    "isResolved": true,
    "parentId": "comment_1",
    "elementId": "shape_1"
  }
]
```

## Additional Notes
- Focus on creating a clean, intuitive interface that makes translation efficient
- Ensure visual hierarchy emphasizes the slide content and editing tools
- Design should feel professional and focused on productivity
- Include appropriate loading and empty states for all data-dependent views
- Ensure sufficient color contrast for text elements
