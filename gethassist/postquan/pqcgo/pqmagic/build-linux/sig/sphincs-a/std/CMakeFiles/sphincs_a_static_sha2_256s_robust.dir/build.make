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
CMAKE_BINARY_DIR = /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux

# Include any dependencies generated for this target.
include sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.make

# Include the progress variables for this target.
include sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/progress.make

# Include the compile flags for this target's objects.
include sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o: ../sig/sphincs-a/std/address.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/address.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/address.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/address.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o: ../sig/sphincs-a/std/fors.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/fors.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/fors.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/fors.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o: ../sig/sphincs-a/std/merkle.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/merkle.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/merkle.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/merkle.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o: ../sig/sphincs-a/std/sign.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/sign.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/sign.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/sign.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o: ../sig/sphincs-a/std/uintx.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/uintx.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/uintx.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/uintx.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o: ../sig/sphincs-a/std/utils.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/utils.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/utils.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/utils.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o: ../sig/sphincs-a/std/utilsx1.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_7) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/utilsx1.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/utilsx1.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/utilsx1.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o: ../sig/sphincs-a/std/wots.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_8) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/wots.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/wots.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/wots.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o: ../sig/sphincs-a/std/wotsx1.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_9) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/wotsx1.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/wotsx1.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/wotsx1.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o: ../sig/sphincs-a/std/hash_sha2.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_10) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/hash_sha2.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/hash_sha2.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/hash_sha2.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o: ../sig/sphincs-a/std/sha2.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_11) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/sha2.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/sha2.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/sha2.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.s

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/flags.make
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o: ../sig/sphincs-a/std/thash_sha2_robust.c
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_12) "Building C object sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o -MF CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o.d -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o -c /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/thash_sha2_robust.c

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.i"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/thash_sha2_robust.c > CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.i

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.s"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/thash_sha2_robust.c -o CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.s

# Object files for target sphincs_a_static_sha2_256s_robust
sphincs_a_static_sha2_256s_robust_OBJECTS = \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o" \
"CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o"

# External object files for target sphincs_a_static_sha2_256s_robust
sphincs_a_static_sha2_256s_robust_EXTERNAL_OBJECTS =

sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/address.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/fors.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/merkle.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sign.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/uintx.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utils.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/utilsx1.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wots.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/wotsx1.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/hash_sha2.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/sha2.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/thash_sha2_robust.c.o
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/build.make
sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a: sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/CMakeFiles --progress-num=$(CMAKE_PROGRESS_13) "Linking C static library libsphincs_a_static_sha2_256s_robust.a"
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && $(CMAKE_COMMAND) -P CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/cmake_clean_target.cmake
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/build: sig/sphincs-a/std/libsphincs_a_static_sha2_256s_robust.a
.PHONY : sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/build

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/clean:
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std && $(CMAKE_COMMAND) -P CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/cmake_clean.cmake
.PHONY : sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/clean

sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/depend:
	cd /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : sig/sphincs-a/std/CMakeFiles/sphincs_a_static_sha2_256s_robust.dir/depend

