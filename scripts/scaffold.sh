#!/bin/bash

# Check if resource name is provided
if [ -z "$1" ]; then
	echo "❌ Resource name required: ./scripts/scaffold.sh task"
	exit 1
fi

# Setup variables
RESOURCE_NAME=$(echo "$1" | sed 's/.*/\u&/g')
resource_name="$1"
HANDLER_DIR="handlers"
MODEL_DIR="models"
HANDLER_FILE="${HANDLER_DIR}/${resource_name}_resource.go"
MODEL_FILE="${MODEL_DIR}/${resource_name}.go"

# Create directories if they don't exist
mkdir -p "$HANDLER_DIR" "$MODEL_DIR"

# Create handler file if it doesn't exist
if [ -f "$HANDLER_FILE" ]; then
	echo "⚠️  $HANDLER_FILE already exists"
else
	echo "🛠 Creating $HANDLER_FILE..."
	sed "s/{{ResourceName}}/$RESOURCE_NAME/g; s/{{resourceName}}/$resource_name/g" templates/handler.tmpl >"$HANDLER_FILE"
fi

# Create model file if it doesn't exist
if [ -f "$MODEL_FILE" ]; then
	echo "⚠️  $MODEL_FILE already exists"
else
	echo "📦 Creating model: $MODEL_FILE"
	sed "s/{{ResourceName}}/$RESOURCE_NAME/g; s/{{resourceName}}/$resource_name/g" templates/model.tmpl >"$MODEL_FILE"
fi

echo "✅ Done. Remember to register route in routes/routes.go:"
echo "    RegisterResourceRoutes(r, \"/${resource_name}s\", handlers.${RESOURCE_NAME}Resource{})"
