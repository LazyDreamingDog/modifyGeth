# Install script for directory: /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std

# Set the install prefix
if(NOT DEFINED CMAKE_INSTALL_PREFIX)
  set(CMAKE_INSTALL_PREFIX "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs")
endif()
string(REGEX REPLACE "/$" "" CMAKE_INSTALL_PREFIX "${CMAKE_INSTALL_PREFIX}")

# Set the install configuration name.
if(NOT DEFINED CMAKE_INSTALL_CONFIG_NAME)
  if(BUILD_TYPE)
    string(REGEX REPLACE "^[^A-Za-z0-9_]+" ""
           CMAKE_INSTALL_CONFIG_NAME "${BUILD_TYPE}")
  else()
    set(CMAKE_INSTALL_CONFIG_NAME "Release")
  endif()
  message(STATUS "Install configuration: \"${CMAKE_INSTALL_CONFIG_NAME}\"")
endif()

# Set the component getting installed.
if(NOT CMAKE_INSTALL_COMPONENT)
  if(COMPONENT)
    message(STATUS "Install component: \"${COMPONENT}\"")
    set(CMAKE_INSTALL_COMPONENT "${COMPONENT}")
  else()
    set(CMAKE_INSTALL_COMPONENT)
  endif()
endif()

# Is this installation the result of a crosscompile?
if(NOT DEFINED CMAKE_CROSSCOMPILING)
  set(CMAKE_CROSSCOMPILING "TRUE")
endif()

# Set default install directory permissions.
if(NOT DEFINED CMAKE_OBJDUMP)
  set(CMAKE_OBJDUMP "/usr/bin/x86_64-w64-mingw32-objdump")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/win/libpqmagic_sphincs_a_std.so")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/win" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/sphincs-a/std/libpqmagic_sphincs_a_std.so")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/win/libpqmagic_sphincs_a_std.a")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/win" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-windows/sig/sphincs-a/std/libpqmagic_sphincs_a_std.a")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/sphincs-a/std/api.h")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/sphincs-a/std" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/api.h")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/sphincs-a/std/params.h")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/sphincs-a/std" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/params.h")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/sphincs-a/std/config.h")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/sphincs-a/std" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/sphincs-a/std/config.h")
endif()

