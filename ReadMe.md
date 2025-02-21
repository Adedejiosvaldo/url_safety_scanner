# URL Safety Scanner

A Go-based web service that provides URL safety analysis using multiple security APIs.

## Features

- 🔍 URL extraction from text messages
- 🛡️ URL safety checking using Google Safe Browsing API
- 🌐 Additional scanning via URLScan.io
- ✨ Real-time URL classification
- 🚦 Safety status indicators
- 🔗 Safe link previews

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/safe_url_scanner.git

# Navigate to project directory
cd safe_url_scanner

# Install dependencies
go mod install

# Set up environment variables
cp .env
```

Environment Variables

```
URL_SCAN_KEY=your_urlscan_io_api_key
GOOGLE_SAFE_BROWSING_API_KEY=your_google_api_key
PORT=8080
```

API Endpoints

1. Scan URL
2. Integration Specification

```HTTP
GET /integration-spec
```

Usage Examples

Scan a URL

```HTTP
POST /scan-url
Content-Type: application/json

{
"message": "Check this link: https://example.com"
}
```

Response Format

```json
{
  "event_name": "url_scanned",
  "message": "✅ URL Check: https://example.com\n→ Status: safe\n→ Recommendation: This link appears safe",
  "urls": ["https://example.com"],
  "status": "success",
  "username": "url-scanner-bot"
}
```

Security Features
URL validation and sanitization
Multiple API security checks
Safe browsing verification
Suspicious link warnings
Hidden URLs for unsafe content
Development

```bash
# Run the server
go run main.go

# Build the application
go build

```

Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request
