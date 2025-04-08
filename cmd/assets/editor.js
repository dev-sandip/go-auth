class MarkdownEditor extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });

    // Load markdown-it and highlight.js from CDN if not already loaded
    this.ensureDependenciesLoaded().then(() => {
      this.md = window.markdownit({
        html: true,
        linkify: true,
        typographer: true,
        highlight: function (str, lang) {
          if (lang && window.hljs && window.hljs.getLanguage(lang)) {
            try {
              return window.hljs.highlight(str, { language: lang }).value;
            } catch (__) { }
          }
          return ''; // use external default escaping
        }
      });
      this.render();
      this.setupEventListeners();
    });
  }

  ensureDependenciesLoaded() {
    return new Promise((resolve) => {
      let dependenciesLoaded = 0;
      const totalDependencies = 2;

      function checkAllLoaded() {
        dependenciesLoaded++;
        if (dependenciesLoaded === totalDependencies) {
          resolve();
        }
      }

      if (!window.markdownit) {
        const markdownScript = document.createElement('script');
        markdownScript.src = 'https://cdnjs.cloudflare.com/ajax/libs/markdown-it/13.0.1/markdown-it.min.js';
        markdownScript.onload = checkAllLoaded;
        document.head.appendChild(markdownScript);
      } else {
        checkAllLoaded();
      }

      if (!window.hljs) {
        // Load highlight.js script
        const hlScript = document.createElement('script');
        hlScript.src = 'https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/highlight.min.js';
        hlScript.onload = () => {
          // Load highlight.js CSS
          const hlCss = document.createElement('link');
          hlCss.rel = 'stylesheet';
          hlCss.href = 'https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/github.min.css';
          document.head.appendChild(hlCss);

          checkAllLoaded();
        };
        document.head.appendChild(hlScript);
      } else {
        checkAllLoaded();
      }
    });
  }

  render() {
    const height = this.getAttribute('height') || '400px';
    const initialValue = this.textContent || '# Hello Markdown\n\nStart typing...\n\n```javascript\n// Code with syntax highlighting\nfunction hello() {\n    console.log("Hello world!");\n}\n```';
    this.textContent = ''; // Clear any existing content

    this.shadowRoot.innerHTML = `
      <style>
        :host {
          display: block;
          margin: 0;
          box-sizing: border-box;
          font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
        }
        
        * {
          box-sizing: border-box;
        }
        
        .container {
          display: flex;
          flex-direction: column;
          height: ${height};
          border: 1px solid #ddd;
          border-radius: 4px;
          overflow: hidden;
        }
        
        .toolbar {
          display: flex;
          gap: 5px;
          padding: 8px;
          background: #f5f5f5;
          border-bottom: 1px solid #ddd;
          flex-wrap: wrap;
        }
        
        .toolbar button {
          background: #fff;
          color: #333;
          border: 1px solid #ccc;
          border-radius: 3px;
          padding: 4px 8px;
          font-size: 14px;
          cursor: pointer;
          margin: 0;
        }
        
        .toolbar button:hover {
          background: #f0f0f0;
        }
        
        .toolbar select {
          height: 28px;
          border: 1px solid #ccc;
          border-radius: 3px;
          background: #fff;
        }
        
        .editor-container {
          display: flex;
          flex: 1;
          min-height: 0;
        }
        
        .editor-section, .preview-section {
          width: 50%;
          flex: 1;
          overflow: auto;
          display: flex;
          flex-direction: column;
        }
        
        .editor-section {
          border-right: 1px solid #ddd;
        }
        
        .section-header {
          padding: 8px;
          background: #eee;
          font-weight: bold;
          border-bottom: 1px solid #ddd;
        }
        
        textarea {
          flex: 1;
          width: 100%;
          padding: 12px;
          border: none;
          resize: none;
          font-family: monospace;
          font-size: 14px;
          line-height: 1.5;
          background: #fff;
          color: #333;
          outline: none;
        }
        
        #preview {
          flex: 1;
          padding: 12px;
          overflow: auto;
          line-height: 1.6;
        }
        
        /* Markdown preview styles */
        #preview h1, #preview h2, #preview h3, #preview h4, #preview h5, #preview h6 {
          margin-top: 16px;
          margin-bottom: 16px;
          line-height: 1.25;
        }
        
        #preview h1 {
          padding-bottom: 0.3em;
          border-bottom: 1px solid #eaecef;
          font-size: 2em;
        }
        
        #preview h2 {
          padding-bottom: 0.3em;
          border-bottom: 1px solid #eaecef;
          font-size: 1.5em;
        }
        
        #preview p {
          margin-bottom: 16px;
        }
        
        #preview pre {
          padding: 16px;
          overflow: auto;
          background-color: #f6f8fa;
          border-radius: 3px;
          margin-bottom: 16px;
        }
        
        #preview code {
          font-family: SFMono-Regular, Consolas, "Liberation Mono", Menlo, monospace;
          background-color: rgba(27, 31, 35, 0.05);
          padding: 0.2em 0.4em;
          border-radius: 3px;
        }
        
        #preview pre code {
          background-color: transparent;
          padding: 0;
        }
        
        #preview blockquote {
          padding: 0 1em;
          color: #6a737d;
          border-left: 0.25em solid #dfe2e5;
          margin-bottom: 16px;
        }
        
        #preview ul, #preview ol {
          padding-left: 2em;
          margin-bottom: 16px;
        }
        
        #preview table {
          border-collapse: collapse;
          margin-bottom: 16px;
          width: 100%;
        }
        
        #preview table th, #preview table td {
          padding: 6px 13px;
          border: 1px solid #dfe2e5;
        }
        
        #preview table tr {
          background-color: #fff;
          border-top: 1px solid #c6cbd1;
        }
        
        #preview table tr:nth-child(2n) {
          background-color: #f6f8fa;
        }
        
        #preview img {
          max-width: 100%;
        }
        
        .status-bar {
          padding: 8px;
          background: #f5f5f5;
          border-top: 1px solid #ddd;
          display: flex;
          justify-content: space-between;
          font-size: 12px;
        }
        
        .export-controls {
          display: flex;
          align-items: center;
          gap: 8px;
        }
        
        /* Make the editor full-width when preview is hidden */
        .no-preview .editor-section {
          width: 100%;
          border-right: none;
        }
        
        .no-preview .preview-section {
          display: none;
        }

        /* Light theme select for syntax highlighting */
        .theme-select {
          margin-left: auto;
        }
      </style>
      
      <div class="container">
        <div class="toolbar">
          <button data-action="header1">H1</button>
          <button data-action="header2">H2</button>
          <button data-action="bold">B</button>
          <button data-action="italic">I</button>
          <button data-action="link">Link</button>
          <button data-action="image">Image</button>
          <button data-action="code">Code</button>
          <button data-action="codeblock">Code Block</button>
          <button data-action="ullist">• List</button>
          <button data-action="ollist">1. List</button>
          <button data-action="quote">Quote</button>
          <button data-action="table">Table</button>
          <div class="theme-select">
            <select id="highlight-theme">
              <option value="github">GitHub</option>
              <option value="github-dark">GitHub Dark</option>
              <option value="monokai">Monokai</option>
              <option value="dracula">Dracula</option>
              <option value="atom-one-dark">Atom One Dark</option>
              <option value="vs">Visual Studio</option>
              <option value="solarized-light">Solarized Light</option>
              <option value="solarized-dark">Solarized Dark</option>
            </select>
          </div>
          <button data-action="togglePreview">Toggle Preview</button>
        </div>
        
        <div class="editor-container">
          <div class="editor-section">
            <div class="section-header">EDITOR</div>
            <textarea id="editor">${initialValue}</textarea>
          </div>
          <div class="preview-section">
            <div class="section-header">PREVIEW</div>
            <div id="preview"></div>
          </div>
        </div>
        
        <div class="status-bar">
          <span id="wordcount">0 words, 0 characters</span>
          <div class="export-controls">
            <select id="export-format">
              <option value="markdown">Markdown</option>
              <option value="html">HTML</option>
            </select>
            <button id="export-btn">Export</button>
          </div>
        </div>
      </div>
    `;

    // Initialize preview
    const editor = this.shadowRoot.getElementById('editor');
    this.updatePreview(editor.value);
    this.updateWordCount(editor.value);
  }

  setupEventListeners() {
    const editor = this.shadowRoot.getElementById('editor');
    const exportBtn = this.shadowRoot.getElementById('export-btn');
    const toolbar = this.shadowRoot.querySelectorAll('.toolbar button');
    const themeSelect = this.shadowRoot.getElementById('highlight-theme');

    // Update preview on input
    editor.addEventListener('input', () => {
      this.updatePreview(editor.value);
      this.updateWordCount(editor.value);
    });

    // Export button click
    exportBtn.addEventListener('click', () => {
      const format = this.shadowRoot.getElementById('export-format').value;
      this.exportContent(format);
    });

    // Toolbar buttons
    toolbar.forEach(button => {
      if (button.getAttribute('data-action') === 'togglePreview') {
        button.addEventListener('click', () => this.togglePreview());
      } else {
        button.addEventListener('click', () => this.insertMarkdown(button.getAttribute('data-action')));
      }
    });

    // Theme change handler
    themeSelect.addEventListener('change', (e) => {
      this.changeHighlightTheme(e.target.value);
    });

    // Tab handling
    editor.addEventListener('keydown', (e) => {
      if (e.key === 'Tab') {
        e.preventDefault();
        const start = editor.selectionStart;
        const end = editor.selectionEnd;

        editor.value = editor.value.substring(0, start) + '    ' + editor.value.substring(end);
        editor.selectionStart = editor.selectionEnd = start + 4;

        this.updatePreview(editor.value);
      }
    });
  }

  updatePreview(markdown) {
    const preview = this.shadowRoot.getElementById('preview');
    preview.innerHTML = this.md.render(markdown);
  }

  updateWordCount(text) {
    const wordcountEl = this.shadowRoot.getElementById('wordcount');
    const words = text.trim() ? text.trim().split(/\s+/).length : 0;
    const chars = text.length;
    wordcountEl.textContent = `${words} words, ${chars} characters`;
  }

  togglePreview() {
    const container = this.shadowRoot.querySelector('.editor-container');
    container.classList.toggle('no-preview');
  }

  changeHighlightTheme(theme) {
    // Remove existing theme CSS if it exists
    const oldThemeLink = document.querySelector('link[data-highlight-theme]');
    if (oldThemeLink) {
      oldThemeLink.remove();
    }

    // Add new theme CSS
    const themeLink = document.createElement('link');
    themeLink.rel = 'stylesheet';
    themeLink.href = `https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/${theme}.min.css`;
    themeLink.setAttribute('data-highlight-theme', theme);
    document.head.appendChild(themeLink);

    // Re-render preview to apply new theme
    const editor = this.shadowRoot.getElementById('editor');
    this.updatePreview(editor.value);
  }

  insertMarkdown(action) {
    const editor = this.shadowRoot.getElementById('editor');
    const start = editor.selectionStart;
    const end = editor.selectionEnd;
    const selectedText = editor.value.substring(start, end);
    let insertText = '';

    switch (action) {
      case 'header1':
        insertText = `# ${selectedText || 'Heading 1'}`;
        break;
      case 'header2':
        insertText = `## ${selectedText || 'Heading 2'}`;
        break;
      case 'bold':
        insertText = `**${selectedText || 'bold text'}**`;
        break;
      case 'italic':
        insertText = `*${selectedText || 'italic text'}*`;
        break;
      case 'link':
        insertText = `[${selectedText || 'link text'}](url)`;
        break;
      case 'image':
        insertText = `![${selectedText || 'alt text'}](image-url)`;
        break;
      case 'code':
        insertText = `\`${selectedText || 'code'}\``;
        break;
      case 'codeblock':
        insertText = `\`\`\`javascript\n${selectedText || '// Code block\nfunction example() {\n    return "Hello world";\n}'}\n\`\`\``;
        break;
      case 'ullist':
        insertText = `- ${selectedText || 'List item'}`;
        break;
      case 'ollist':
        insertText = `1. ${selectedText || 'List item'}`;
        break;
      case 'quote':
        insertText = `> ${selectedText || 'Blockquote'}`;
        break;
      case 'table':
        insertText = `| Header 1 | Header 2 |\n| -------- | -------- |\n| Cell 1   | Cell 2   |`;
        break;
    }

    // Insert text
    editor.focus();
    editor.value = editor.value.substring(0, start) + insertText + editor.value.substring(end);
    editor.selectionStart = start + insertText.length;
    editor.selectionEnd = start + insertText.length;

    // Update preview
    this.updatePreview(editor.value);
    this.updateWordCount(editor.value);
  }

  exportContent(format) {
    const editor = this.shadowRoot.getElementById('editor');
    let content, mimeType, filename;

    if (format === 'markdown') {
      content = editor.value;
      mimeType = 'text/markdown';
      filename = 'document.md';
    } else {
      // For HTML export, include basic styling and the current highlight.js theme
      const htmlContent = this.md.render(editor.value);
      const theme = this.shadowRoot.getElementById('highlight-theme').value;

      content = `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Exported Markdown</title>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/${theme}.min.css">
  <style>
    body {
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
      line-height: 1.6;
      max-width: 800px;
      margin: 0 auto;
      padding: 20px;
    }
    pre {
      background: #f6f8fa;
      padding: 16px;
      overflow: auto;
      border-radius: 3px;
    }
    code {
      font-family: SFMono-Regular, Consolas, "Liberation Mono", Menlo, monospace;
    }
    blockquote {
      border-left: 4px solid #dfe2e5;
      padding-left: 16px;
      margin-left: 0;
      color: #6a737d;
    }
    table {
      border-collapse: collapse;
      width: 100%;
    }
    table th, table td {
      border: 1px solid #dfe2e5;
      padding: 6px 13px;
    }
  </style>
</head>
<body>
  ${htmlContent}
</body>
</html>`;
      mimeType = 'text/html';
      filename = 'document.html';
    }

    // Create download link
    const blob = new Blob([content], { type: mimeType });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = filename;
    document.body.appendChild(a);
    a.click();

    // Clean up
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  }
}

// Register the custom element
customElements.define('markdown-editor', MarkdownEditor);