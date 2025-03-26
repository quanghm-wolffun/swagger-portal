# OpenAPI Workflow Documentation

This document outlines a streamlined workflow for developing and sharing OpenAPI 3.0 specifications.

## Overview

The workflow consists of five key steps:

1. **Write API spec in YAML** (manual)
2. **Convert to API spec in JSON** (automated)
3. **Format spec files** (automated)
4. **Render spec locally** (automated)
5. **Push to company host** (automated)

Only step 1 requires manual work. The other steps can be automated using the tools described below.

## Prerequisites

- Go (for running `yq`)
- Bun (for running JavaScript tools)
- Git (for version control)

## Tools

- [yq](https://github.com/mikefarah/yq) - YAML processor
- [Prettier](https://prettier.io/) - Code formatter
- [Redocly CLI](https://redocly.com/docs/cli/) - API documentation renderer

## Step-by-Step Workflow

### 1. Write API Specification in YAML

Create or edit your API specification in YAML format at `./docs/swagger.yaml`:

```yaml
openapi: 3.0.0
info:
  title: Example API
  version: 1.0.0
# Continue defining your API...
```

### 2. Convert YAML to JSON

Convert your YAML specification to JSON format:

```bash
# Install yq if not already installed
go install github.com/mikefarah/yq/v4@latest

# Convert YAML to JSON using our converter
go run convert.go
```

### 3. Format Specification Files

Format both YAML and JSON files for consistency:

```bash
# Format using Prettier
bunx prettier --write ./docs/swagger.yaml ./docs/swagger.json
```

### 4. Validate the OpenAPI Specification

Validate your OpenAPI specification before previewing:

```bash
# Validate using Redocly CLI
bunx @redocly/cli lint ./docs/swagger.json
```

### 5. Preview Documentation Locally

Render your API documentation locally to review before sharing:

```bash
# Install Redocly CLI if needed
bun add @redocly/cli

# Preview documentation using Redocly CLI
bunx @redocly/cli preview-docs ./docs/swagger.json
```

This will start a local server and open your browser to view the rendered documentation.

### 6. Share with Other Teams

Push your specification to the company host for other teams to access:

```bash
# Example - adjust according to your company's hosting solution
git add ./docs/swagger.yaml ./docs/swagger.json
git commit -m "Update API specification"
git push origin main
```

## Automation Script

To automate steps 2–5, use the npm scripts provided:

```bash
# Run all steps
bun run update-docs
```

Or use them individually:

```bash
# Convert YAML to JSON
bun run convert

# Format files
bun run format

# Validate specification
bun run validate

# Preview documentation
bun run preview
```

## Troubleshooting

### JSON Parsing Errors

If you encounter JSON parsing errors with Redocly, check for:

1. **Encoding issues**: Ensure your JSON file is properly encoded without special characters
2. **File corruption**: Regenerate the JSON from YAML using `go run convert.go`
3. **Valid OpenAPI format**: Validate your OpenAPI specification using our validation script:
   ```bash
   bun run validate
   ```

### PowerShell Command Issues

When running commands in PowerShell, you might see error messages if commands are not properly formatted:

1. Make sure to use `bunx` instead of `npx` when using Bun
2. Use `bun run <script-name>` to execute npm scripts
3. If encountering "Failed to parse API description" errors, try running the validation step first

### Redocly Community Edition Notice

When using the community edition of Redocly, you'll see a notice:
```
Using Redoc community edition.
Login with redocly login or use an enterprise license key to preview with the premium docs.
```

This is normal and doesn't affect basic functionality. The free community edition is sufficient for most needs.

## Best Practices

1. Always edit the YAML version as the source of truth
2. Run the conversion and formatting before committing changes
3. Always validate your OpenAPI specification before previewing
4. Preview locally to catch any issues before sharing
5. Use version control to track changes to your API specification
