#!/bin/bash

# End execution on error
set -e

echo "Generate High lvl activity diag using Graphviz"
dot -Tpng pixiActivity.dot -o out/pixiActivity.png


mkdir -p lib
echo "Download plantuml.jar if it doesn't exist as ./lib/plantuml.jar"
if ! sha256sum -c plantuml.sha256   ; then
    wget "http://sourceforge.net/projects/plantuml/files/plantuml.1.2019.1.jar" -O lib/plantuml.jar

fi
echo "Generate diagram : domain model"
java -jar lib/plantuml.jar  DiagClassSystem.uml -o out/
