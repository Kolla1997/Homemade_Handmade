
#!/bin/bash
echo "🚀 Building frontend..."
npm run build
echo "🚀 Starting Go backend server..."
cd server-go
go run .
