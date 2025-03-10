#!/bin/bash
debug=false

# base_Path="/tmp/slicer-cli"
base_Path="$(realpath $PWD/..)"

source ../golang.env

if [ $debug = true ]; then
	echo -e "$BLUE🏷️ base_Path: $YELLOW$base_Path$NC"
	echo -e "$BLUE🏷️ buildpath: $YELLOW$buildpath$NC"
	echo -e "$BLUE🏷️ buildpath_slicer: $YELLOW$buildpath_slicer$ND"
fi

if [ ! -d ${buildpath_slicer} ]; then
	echo -e "$GREEN📁Creating build directory$NC"
	mkdir -p ${buildpath_slicer}
	if [ ! -d ${buildpath_slicer} ]; then
		echo -e "$RED❌Failed to create build directory: $YELLOW${buildpath_slicer}$NC"
		return 1
	fi
elif [ $debug = true ]; then
	echo -e "$YELLOW📁Build directory already exists$NC"
fi

echo -e "$GREEN🚧Clean-Up previous built files$NC"
rm -rdf ${buildpath_slicer}/*

echo -e "$GREEN🚧Building slicer-cli$NC"
go build -o ${buildpath_slicer}/$binname_slicer -v ${PWD}

if [ ! -f ${buildpath_slicer}/$binname_slicer ]; then
	echo -e "\t$RED❌Failed to build slicer-cli$NC"
	return 1
else
	echo -e "\t$GREEN✅slicer-cli built successfully$NC"
	echo -e "\t$BLUE🏷️Executable: $YELLOW${buildpath_slicer}/$binname_slicer$NC"
fi
