clean:
	@echo "🧹 Cleaning up empty GO-files..."
	@finf . -type f -name "*.go" -exec bash -c '[[ -s "{}" ]] || echo "Deleting empty file: {}" && rm -f "{}"' \;
