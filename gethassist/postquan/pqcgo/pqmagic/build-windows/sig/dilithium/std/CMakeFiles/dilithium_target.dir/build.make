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

# Utility rule file for dilithium_target.

# Include any custom commands dependencies for this target.
include sig/dilithium/std/CMakeFiles/dilithium_target.dir/compiler_depend.make

# Include the progress variables for this target.
include sig/dilithium/std/CMakeFiles/dilithium_target.dir/progress.make

sig/dilithium/std/CMakeFiles/dilithium_target: sig/dilithium/std/libdilithium_static_2.a
sig/dilithium/std/CMakeFiles/dilithium_target: sig/dilithium/std/libdilithium_static_3.a
sig/dilithium/std/CMakeFiles/dilithium_target: sig/dilithium/std/libdilithium_static_5.a
sig/dilithium/std/CMakeFiles/dilithium_target: hash/keccak/libfips202.a
sig/dilithium/std/CMakeFiles/dilithium_target: utils/librandombytes.a
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --blue --bold --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Generating shared library libpqmagic_dilithium_std.so"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std && /usr/bin/x86_64-w64-mingw32-gcc -shared -o /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std/libpqmagic_dilithium_std.so -Wl,--whole-archive /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std/libdilithium_static_2.a /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std/libdilithium_static_3.a /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std/libdilithium_static_5.a /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/hash/keccak/libfips202.a /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/utils/librandombytes.a -Wl,--no-whole-archive

dilithium_target: sig/dilithium/std/CMakeFiles/dilithium_target
dilithium_target: sig/dilithium/std/CMakeFiles/dilithium_target.dir/build.make
.PHONY : dilithium_target

# Rule to build all files generated by this target.
sig/dilithium/std/CMakeFiles/dilithium_target.dir/build: dilithium_target
.PHONY : sig/dilithium/std/CMakeFiles/dilithium_target.dir/build

sig/dilithium/std/CMakeFiles/dilithium_target.dir/clean:
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std && $(CMAKE_COMMAND) -P CMakeFiles/dilithium_target.dir/cmake_clean.cmake
.PHONY : sig/dilithium/std/CMakeFiles/dilithium_target.dir/clean

sig/dilithium/std/CMakeFiles/dilithium_target.dir/depend:
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/dilithium/std /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/dilithium/std/CMakeFiles/dilithium_target.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : sig/dilithium/std/CMakeFiles/dilithium_target.dir/depend

