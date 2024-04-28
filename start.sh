#!/bin/sh
echo "Start"

echo $GOLANG_ENVIRONMENT
echo $API_VERSION

python3 -u rogue.py 'appsettings.Template.json' "$PROJECT_CONFIG" $GOLANG_ENVIRONMENT
./main $@