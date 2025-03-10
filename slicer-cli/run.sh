#!/bin/bash

#!/bin/bash
debug=true

# base_Path="/tmp/slicer-cli"
base_Path="$(realpath $PWD/..)"

. build.sh

echo -e "🪄Running slicer-cli"

if [ ! -f ${buildpath_slicer}/$binname_slicer ]; then
	echo -e "\t$RED❌slicer-cli not found$NC"
	return 1
else
	${buildpath_slicer}/$binname_slicer
fi
