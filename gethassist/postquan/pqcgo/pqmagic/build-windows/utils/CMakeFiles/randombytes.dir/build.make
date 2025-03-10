# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.22

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows

# Include any dependencies generated for this target.
include utils/CMakeFiles/randombytes.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include utils/CMakeFiles/randombytes.dir/compiler_depend.make

# Include the progress variables for this target.
include utils/CMakeFiles/randombytes.dir/progress.make

# Include the compile flags for this target's objects.
include utils/CMakeFiles/randombytes.dir/flags.make

utils/CMakeFiles/randombytes.dir/randombytes.c.obj: utils/CMakeFiles/randombytes.dir/flags.make
utils/CMakeFiles/randombytes.dir/randombytes.c.obj: utils/CMakeFiles/randombytes.dir/includes_C.rsp
utils/CMakeFiles/randombytes.dir/randombytes.c.obj: ../utils/randombytes.c
utils/CMakeFiles/randombytes.dir/randombytes.c.obj: utils/CMakeFiles/randombytes.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object utils/CMakeFiles/randombytes.dir/randombytes.c.obj"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils && /usr/bin/x86_64-w64-mingw32-gcc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT utils/CMakeFiles/randombytes.dir/randombytes.c.obj -MF CMakeFiles/randombytes.dir/randombytes.c.obj.d -o CMakeFiles/randombytes.dir/randombytes.c.obj -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/utils/randombytes.c

utils/CMakeFiles/randombytes.dir/randombytes.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/randombytes.dir/randombytes.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils && /usr/bin/x86_64-w64-mingw32-gcc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/utils/randombytes.c > CMakeFiles/randombytes.dir/randombytes.c.i

utils/CMakeFiles/randombytes.dir/randombytes.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/randombytes.dir/randombytes.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils && /usr/bin/x86_64-w64-mingw32-gcc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/utils/randombytes.c -o CMakeFiles/randombytes.dir/randombytes.c.s

# Object files for target randombytes
randombytes_OBJECTS = \
"CMakeFiles/randombytes.dir/randombytes.c.obj"

# External object files for target randombytes
randombytes_EXTERNAL_OBJECTS =

utils/librandombytes.a: utils/CMakeFiles/randombytes.dir/randombytes.c.obj
utils/librandombytes.a: utils/CMakeFiles/randombytes.dir/build.make
utils/librandombytes.a: utils/CMakeFiles/randombytes.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C static library librandombytes.a"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils && $(CMAKE_COMMAND) -P CMakeFiles/randombytes.dir/cmake_clean_target.cmake
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/randombytes.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
utils/CMakeFiles/randombytes.dir/build: utils/librandombytes.a
.PHONY : utils/CMakeFiles/randombytes.dir/build

utils/CMakeFiles/randombytes.dir/clean:
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils && $(CMAKE_COMMAND) -P CMakeFiles/randombytes.dir/cmake_clean.cmake
.PHONY : utils/CMakeFiles/randombytes.dir/clean

utils/CMakeFiles/randombytes.dir/depend:
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/utils /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils/CMakeFiles/randombytes.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : utils/CMakeFiles/randombytes.dir/depend

