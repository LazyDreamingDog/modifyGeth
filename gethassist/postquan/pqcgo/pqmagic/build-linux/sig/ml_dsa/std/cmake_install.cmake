# Install script for directory: /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/ml_dsa/std

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

# Install shared libraries without execute permission?
if(NOT DEFINED CMAKE_INSTALL_SO_NO_EXE)
  set(CMAKE_INSTALL_SO_NO_EXE "1")
endif()

# Is this installation the result of a crosscompile?
if(NOT DEFINED CMAKE_CROSSCOMPILING)
  set(CMAKE_CROSSCOMPILING "FALSE")
endif()

# Set default install directory permissions.
if(NOT DEFINED CMAKE_OBJDUMP)
  set(CMAKE_OBJDUMP "/usr/bin/objdump")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/linux/libpqmagic_ml_dsa_std.so")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/linux" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/ml_dsa/std/libpqmagic_ml_dsa_std.so")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/linux/libpqmagic_ml_dsa_std.a")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib/linux" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/build-linux/sig/ml_dsa/std/libpqmagic_ml_dsa_std.a")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/ml_dsa/api.h")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/ml_dsa" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/ml_dsa/std/api.h")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/ml_dsa/params.h")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/ml_dsa" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/ml_dsa/std/params.h")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  list(APPEND CMAKE_ABSOLUTE_DESTINATION_FILES
   "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/ml_dsa/config.h")
  if(CMAKE_WARN_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(WARNING "ABSOLUTE path INSTALL DESTINATION : ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  if(CMAKE_ERROR_ON_ABSOLUTE_INSTALL_DESTINATION)
    message(FATAL_ERROR "ABSOLUTE path INSTALL DESTINATION forbidden (by caller): ${CMAKE_ABSOLUTE_DESTINATION_FILES}")
  endif()
  file(INSTALL DESTINATION "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include/sig/ml_dsa" TYPE FILE FILES "/home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/pqmagic/sig/ml_dsa/std/config.h")
endif()

