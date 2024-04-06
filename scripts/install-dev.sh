cd ..

# Build binary
go build -ldflags "-s -w"

# Install binary
mv ./cl /usr/bin
chmod +x /usr/bin/cl

# Move to back to scripts directory
cd scripts