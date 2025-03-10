cmake_minimum_required(VERSION 3.10)

# PQMagic version.
project(PQMagic)

# Set c standard
set(CMAKE_BUILD_TYPE Release)
set(CMAKE_C_STANDARD 11)

# Select PQMagic Version.
# Open Source Version Only Support PQMagic-std.
set(PQMAGIC_VERSION "std" CACHE STRING "Set PQMagic version")

if(NOT PQMAGIC_VERSION STREQUAL "std")
    message(FATAL_ERROR "Open Source Version Only Support PQMagic-std. Please set PQMAGIC_VERSION as \"std\" or contact as for further high performance support.")
endif()
message(STATUS "Use PQMagic-${PQMAGIC_VERSION}")

# Set default algorithim mode.
## KEM
set(DEFAULT_KYBER_MODES 2 3 4 CACHE STRING "All modes for Kyber algorithm")
set(DEFAULT_ML_KEM_MODES 512 768 1024 CACHE STRING "All modes for ML-KEM algorithm (FIPS 203)")
set(DEFAULT_AIGIS_ENC_MODES 1 2 3 4 CACHE STRING "All modes for Aigis-enc algorithm")
## SIG
set(DEFAULT_DILITHIUM_MODES 2 3 5 CACHE STRING "All modes for Dilithium algorithm")
set(DEFAULT_ML_DSA_MODES 44 65 87 CACHE STRING "All modes for ML-DSA algorithm (FIPS 204)")
set(DEFAULT_AIGIS_SIG_MODES 1 2 3 CACHE STRING "All modes for Aigis-sig algorithm")
set(DEFAULT_SLH_DSA_MODES 
    slh-dsa-sm3-128f slh-dsa-sm3-128s 
    slh-dsa-sha2-128f slh-dsa-sha2-128s slh-dsa-sha2-192f slh-dsa-sha2-192s slh-dsa-sha2-256f slh-dsa-sha2-256s 
    slh-dsa-shake-128f slh-dsa-shake-128s slh-dsa-shake-192f slh-dsa-shake-192s slh-dsa-shake-256f slh-dsa-shake-256s 
    CACHE STRING "All modes for SLH-DSA (FIPS 205 / SPHINCS+) algorithm")
set(DEFAULT_SPHINCS_A_MODES 
    sphincs-a-sm3-128f sphincs-a-sm3-128s 
    sphincs-a-sha2-128f sphincs-a-sha2-128s sphincs-a-sha2-192f sphincs-a-sha2-192s sphincs-a-sha2-256f sphincs-a-sha2-256s 
    sphincs-a-shake-128f sphincs-a-shake-128s sphincs-a-shake-192f sphincs-a-shake-192s sphincs-a-shake-256f sphincs-a-shake-256s 
    CACHE STRING "All modes for SPHINCS-Alpha algorithm")
set(DEFAULT_THASH robust simple CACHE STRING "All thash modes for SLH-DSA (FIPS 205 / SPHINCS+) and SPHINCS-Alpha algorithm") # Could be robust and simple

# User can pass mode for alg mode.
## KEM
set(KYBER_MODES ${DEFAULT_KYBER_MODES} CACHE STRING "Selected mode(s) for Kyber algorithm")
set(ML_KEM_MODES ${DEFAULT_ML_KEM_MODES} CACHE STRING "Selected mode(s) for ML-KEM algorithm (FIPS 203)")
set(AIGIS_ENC_MODES ${DEFAULT_AIGIS_ENC_MODES} CACHE STRING "Selected mode(s) for Aigis-enc algorithm")
## SIG
set(DILITHIUM_MODES ${DEFAULT_DILITHIUM_MODES} CACHE STRING "Selected mode(s) for Dilithium algorithm")
set(ML_DSA_MODES ${DEFAULT_ML_DSA_MODES} CACHE STRING "Selected mode(s) for ML-DSA algorithm (FIPS 204)")
set(AIGIS_SIG_MODES ${DEFAULT_AIGIS_SIG_MODES} CACHE STRING "Selected mode(s) for Aigis-sig algorithm")
set(SLH_DSA_MODES ${DEFAULT_SLH_DSA_MODES} CACHE STRING "Selected mode(s) for SLH-DSA (FIPS 205 / SPHINCS+) algorithm")
set(SLH_DSA_THASH ${DEFAULT_THASH} CACHE STRING "Selected thash mode(s) for SLH-DSA (FIPS 205 / SPHINCS+) algorithm")
set(SPHINCS_A_MODES ${DEFAULT_SPHINCS_A_MODES} CACHE STRING "Selected mode(s) for SPHINCS-Alpha algorithm")
set(SPHINCS_A_THASH ${DEFAULT_THASH} CACHE STRING "Selected thash mode(s) for SPHINCS-Alpha algorithm")

## KEM
option(ENABLE_KYBER "Enable Kyber" ON)
option(ENABLE_ML_KEM "Enable ML-KEM" ON)
option(ENABLE_AIGIS_ENC "Enable Aigis-enc" ON)
## SIG
option(ENABLE_DILITHIUM "Enable Dilithium" ON)
option(ENABLE_ML_DSA "Enable ML-DSA" ON)
option(ENABLE_AIGIS_SIG "Enable Aigis-sig" ON)
option(ENABLE_SLH_DSA "Enable SLH-DSA" ON)
option(ENABLE_SPHINCS_A "Enable SPHINCS-Alpha" ON)

if(
    NOT ENABLE_KYBER AND 
    NOT ENABLE_ML_KEM AND
    NOT ENABLE_AIGIS_ENC AND
    NOT ENABLE_DILITHIUM AND 
    NOT ENABLE_ML_DSA AND
    NOT ENABLE_AIGIS_SIG AND
    NOT ENABLE_SLH_DSA AND
    NOT ENABLE_SPHINCS_A)

    message(FATAL_ERROR "At least enalble one algorithm, or use default config with all algorithm enabled. \nCurrent option: ENABLE_ML_KEM: ${ENABLE_ML_KEM}, ENABLE_KYBER: ${ENABLE_KYBER}, ENABLE_AIGIS_ENC: ${ENABLE_AIGIS_ENC}, ENABLE_ML_DSA: ${ENABLE_ML_DSA}, ENABLE_SLH_DSA: ${ENABLE_SLH_DSA}, ENABLE_DILITHIUM: ${ENABLE_DILITHIUM}, ENABLE_AIGIS_SIG: ${ENABLE_AIGIS_SIG}, ENABLE_SPHINCS_A: ${ENABLE_SPHINCS_A}. \nSupport algorithm: ML-KEM, Kyber, Aigis-enc, ML-DSA, SLH-DSA, Dilithium, Aigis-sig, SPHINCS-Alpha.")
endif()

## Hash Function
#  Default use sm3.
#  This hash selector do not work for SLH_DSA and SPHINCS_A
#  Change SLH_DSA_MODES and SPHINCS_A_MODES to select hash mode of SLH_DSA and SPHINCS_A.
option(USE_SM3 "Use sm3 as hash component." ON)
option(USE_SHAKE "Use shake as hash component." OFF)

if(USE_SHAKE)
    set(USE_SM3 OFF)
endif()

if( 
    (USE_SHAKE AND USE_SM3) OR 
    ((NOT USE_SHAKE) AND (NOT USE_SM3)))
    message(FATAL_ERROR "Choose hash mode by -DUSE_SM3=ON or -DUSE_SHAKE=ON")
endif()

if(USE_SM3)
message(STATUS "Hash Component: SM3")
elseif(USE_SHAKE)
message(STATUS "Hash Component: SHAKE")
endif()

# Set global compile options
if(NOT CMAKE_C_COMPILER_ID STREQUAL "Emscripten")
    if(NOT WIN32)
        add_compile_options(
                -Wall -Wextra -Wpedantic -Wmissing-prototypes -Wredundant-decls -Wshadow -fPIC
                -Wpointer-arith -Wa,--noexecstack -fomit-frame-pointer -O3 -mtune=native
        )
    else()
        add_compile_options(
                -Wall -Wextra -Wpedantic -Wmissing-prototypes -Wredundant-decls -Wshadow -fPIC
                -Wpointer-arith -fomit-frame-pointer -O3 -mtune=native
        )
    endif()
else()
    add_compile_options(
            -Wall -Wextra -Wpedantic -Wmissing-prototypes -Wredundant-decls -Wshadow -fPIC
            -Wpointer-arith -Wa,--noexecstack -fomit-frame-pointer -O3
    )
endif()

#################################
# Project library build setting #
#################################
# First get install dir.
set(INSTALL_DIR ${CMAKE_INSTALL_PREFIX} CACHE PATH "Installation directory")
# message(STATUS "Installation directory for PQMagic is: ${INSTALL_DIR}")
# Set install path
set(INSTALL_BIN_DIR ${INSTALL_DIR}/bin CACHE PATH "Executable installation directory")
# 添加跨平台编译支持
if(DEFINED CMAKE_TOOLCHAIN_FILE)
    set(INSTALL_LIB_DIR ${INSTALL_DIR}/lib/win CACHE PATH "Library installation directory")
else ()
    set(INSTALL_LIB_DIR ${INSTALL_DIR}/lib/linux CACHE PATH "Library installation directory")
endif ()
set(INSTALL_INCLUDE_DIR ${INSTALL_DIR}/include CACHE PATH "Include installation directory")

# Show abs path for user.
get_filename_component(INSTALL_BIN_DIR_ABS ${INSTALL_BIN_DIR} ABSOLUTE)
get_filename_component(INSTALL_LIB_DIR_ABS ${INSTALL_LIB_DIR} ABSOLUTE)
get_filename_component(INSTALL_INCLUDE_DIR_ABS ${INSTALL_INCLUDE_DIR} ABSOLUTE)
message(STATUS "Installation executable path: ${INSTALL_BIN_DIR_ABS}")
message(STATUS "Installation library path: ${INSTALL_LIB_DIR_ABS}")
message(STATUS "Installation include path: ${INSTALL_INCLUDE_DIR_ABS}")

# Add subdir.
set(SUPPORT_ALG_STATIC_TARGET "")
set(SUPPORT_ALG_STATIC_TARGET_PATH "")

# 添加跨平台编译支持
if(DEFINED CMAKE_TOOLCHAIN_FILE)
    set(PQMAGIC_SHARED_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.dll")
    set(PQMAGIC_STATIC_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.lib")
    set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR})
else()
    set(PQMAGIC_SHARED_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.so")
    set(PQMAGIC_STATIC_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.a")
endif()

add_subdirectory(hash/sm3) # Add sm3
add_subdirectory(hash/keccak) # Add fips202
add_subdirectory(utils)
## KEM
if(ENABLE_KYBER)
    message(STATUS "KYBER: ENABLED with mode ${KYBER_MODES}")
    add_subdirectory(kem/kyber/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET kyber_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/kem/kyber/${PQMAGIC_VERSION}/libpqmagic_kyber_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "KYBER: DISABLED")
endif()
if(ENABLE_ML_KEM)
    message(STATUS "ML-KEM (FIPS 203): ENABLED with mode ${ML_KEM_MODES}")
    add_subdirectory(kem/ml_kem/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET ml_kem_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/kem/ml_kem/${PQMAGIC_VERSION}/libpqmagic_ml_kem_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "ML-KEM (FIPS 203): DISABLED")
endif()
if(ENABLE_AIGIS_ENC)
    message(STATUS "AIGIS-ENC: ENABLED with mode ${AIGIS_ENC_MODES}")
    add_subdirectory(kem/aigis-enc/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET aigis_enc_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/kem/aigis-enc/${PQMAGIC_VERSION}/libpqmagic_aigis_enc_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "AIGIS-ENC: DISABLED")
endif()
## SIG
if(ENABLE_DILITHIUM)
    message(STATUS "DILITHIUM: ENABLED with mode ${DILITHIUM_MODES}")
    add_subdirectory(sig/dilithium/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET dilithium_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/sig/dilithium/${PQMAGIC_VERSION}/libpqmagic_dilithium_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "DILITHIUM: DISABLED")
endif()
if(ENABLE_ML_DSA)
    message(STATUS "ML-DSA (FIPS 204): ENABLED with mode ${ML_DSA_MODES}")
    add_subdirectory(sig/ml_dsa/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET ml_dsa_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/sig/ml_dsa/${PQMAGIC_VERSION}/libpqmagic_ml_dsa_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "ML-DSA (FIPS 204): DISABLED")
endif()
if(ENABLE_AIGIS_SIG)
    message(STATUS "AIGIS_SIG: ENABLED with mode ${AIGIS_SIG_MODES}")
    add_subdirectory(sig/aigis-sig/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET aigis_sig_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/sig/aigis-sig/${PQMAGIC_VERSION}/libpqmagic_aigis_sig_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "AIGIS_SIG: DISABLED")
endif()
if(ENABLE_SLH_DSA)
    message(STATUS "SLH-DSA (FIPS 205): ENABLED with mode ${SLH_DSA_MODES}")
    add_subdirectory(sig/slh_dsa/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET slh_dsa_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/sig/slh_dsa/${PQMAGIC_VERSION}/libpqmagic_slh_dsa_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "SLH_DSA (FIPS 205): DISABLED")
endif()
if(ENABLE_SPHINCS_A)
    message(STATUS "SPHINCS-Alpha: ENABLED with mode ${SPHINCS_A_MODES}")
    add_subdirectory(sig/sphincs-a/${PQMAGIC_VERSION})
    # Record target name and path.
    list(APPEND SUPPORT_ALG_STATIC_TARGET sphincs_a_static_target)
    list(APPEND SUPPORT_ALG_STATIC_TARGET_PATH "${CMAKE_BINARY_DIR}/sig/sphincs-a/${PQMAGIC_VERSION}/libpqmagic_sphincs_a_${PQMAGIC_VERSION}.a")
else()
    message(STATUS "SPHINCS-Alpha: DISABLED")
endif()

# Common include dir.
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/include)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/utils)
include_directories(${CMAKE_CURRENT_SOURCE_DIR})


# Set a custom target for shared library
add_custom_target(pqmagic_target ALL
    COMMAND ${CMAKE_C_COMPILER} -shared -o ${PQMAGIC_SHARED_LIB_NAME}
            -Wl,--whole-archive 
            ${SUPPORT_ALG_STATIC_TARGET_PATH}
            $<TARGET_FILE:sm3>
            $<TARGET_FILE:fips202>
            $<TARGET_FILE:randombytes>
            -Wl,--no-whole-archive
    DEPENDS ${SUPPORT_ALG_STATIC_TARGET} sm3 fips202 randombytes
    COMMENT "Generating shared library ${PQMAGIC_SHARED_LIB_NAME}"
)

# sm3 and randombytes target's default path.
set(SM3_LIB_PATH ${CMAKE_BINARY_DIR}/hash/sm3)
set(FIPS202_LIB_PATH ${CMAKE_BINARY_DIR}/hash/keccak)
set(RANDOMBYTES_LIB_PATH ${CMAKE_BINARY_DIR}/utils)
# Set mri script for archive library combination.
file(WRITE ${CMAKE_BINARY_DIR}/ar_script.mri "CREATE ${PQMAGIC_STATIC_LIB_NAME}\n")
foreach(lib_path ${SUPPORT_ALG_STATIC_TARGET_PATH})
    file(APPEND ${CMAKE_BINARY_DIR}/ar_script.mri "ADDLIB ${lib_path}\n")
endforeach()
file(APPEND ${CMAKE_BINARY_DIR}/ar_script.mri "ADDLIB ${SM3_LIB_PATH}/libsm3.a\n")
file(APPEND ${CMAKE_BINARY_DIR}/ar_script.mri "ADDLIB ${FIPS202_LIB_PATH}/libfips202.a\n")
file(APPEND ${CMAKE_BINARY_DIR}/ar_script.mri "ADDLIB ${RANDOMBYTES_LIB_PATH}/librandombytes.a\n")
file(APPEND ${CMAKE_BINARY_DIR}/ar_script.mri "SAVE\nEND\n")

# Set a custom target for static library
add_custom_target(pqmagic_static_target ALL
    COMMAND ${CMAKE_AR} -M < ${CMAKE_BINARY_DIR}/ar_script.mri
    DEPENDS ${SUPPORT_ALG_STATIC_TARGET} sm3 fips202 randombytes
    COMMENT "Generating static library CREATE ${PQMAGIC_STATIC_LIB_NAME}\n"
)

# Add cross-compilation for Windows
if(DEFINED CMAKE_TOOLCHAIN_FILE)
    set(PQMAGIC_SHARED_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.dll")
    set(PQMAGIC_STATIC_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.lib")
else()
    set(PQMAGIC_SHARED_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.so")
    set(PQMAGIC_STATIC_LIB_NAME "libpqmagic_${PQMAGIC_VERSION}.a")
endif()

# Install lib.
install(FILES ${CMAKE_BINARY_DIR}/${PQMAGIC_SHARED_LIB_NAME}
    DESTINATION ${INSTALL_LIB_DIR})
install(FILES ${CMAKE_BINARY_DIR}/${PQMAGIC_STATIC_LIB_NAME}
    DESTINATION ${INSTALL_LIB_DIR})
# Custom install all executable target.
install(CODE 
    "
    message(STATUS \"Create soft link library for: ${INSTALL_LIB_DIR}/${PQMAGIC_STATIC_LIB_NAME}\")
    execute_process(COMMAND ln -s ${INSTALL_LIB_DIR}/${PQMAGIC_STATIC_LIB_NAME} ${INSTALL_LIB_DIR}/libpqmagic.a)
    message(STATUS \"Create soft link library for: ${INSTALL_LIB_DIR}/${PQMAGIC_SHARED_LIB_NAME}\")
    execute_process(COMMAND ln -s ${INSTALL_LIB_DIR}/${PQMAGIC_SHARED_LIB_NAME} ${INSTALL_LIB_DIR}/libpqmagic.so)
    "
)

###########################
# Add test binary for alg #
###########################
# Used for add test correctness binary
function(add_algorithm_test TARGET_NAME SOURCE_FILE ALG_NAME TYPE MODE)

    # Define other argument.
    set(options)
    set(oneValueArgs THASH MODE_NAMESPACE)
    set(multiValueArgs)

    # Parse argumanets
    cmake_parse_arguments(ADD_ALGO_TEST "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})

    if(NOT ADD_ALGO_TEST_THASH)
        set(ADD_ALGO_TEST_THASH "")
    endif()
    if(NOT ADD_ALGO_TEST_MODE_NAMESPACE)
        set(ADD_ALGO_TEST_MODE_NAMESPACE "")
    endif()

    add_executable(${TARGET_NAME} ${SOURCE_FILE})
    target_link_libraries(${TARGET_NAME} PRIVATE pqmagic)  # use shared lib.
    add_dependencies(${TARGET_NAME} pqmagic_target)
    set_target_properties(${TARGET_NAME} PROPERTIES
        RUNTIME_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}"
    )
    # Include dir for using lib.
    target_include_directories(${TARGET_NAME} PRIVATE 
        ${CMAKE_CURRENT_SOURCE_DIR}/${TYPE}/${ALG_NAME}/${PQMAGIC_VERSION}
    )

    if (ALG_NAME STREQUAL "aigis-sig")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            AIGIS_SIG_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "dilithium")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            DILITHIUM_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "ml_dsa")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            ML_DSA_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "kyber")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            KYBER_K=${MODE}
        )
    elseif (ALG_NAME STREQUAL "aigis-enc")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            AIGIS_ENC_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "ml_kem")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            ML_KEM_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "slh_dsa")
        target_compile_definitions(
            ${TARGET_NAME} PRIVATE
            SLH_DSA_MODE=${MODE} 
            THASH=${ADD_ALGO_TEST_THASH}
            SLH_DSA_HASH_MODE_NAMESPACE=${ADD_ALGO_TEST_MODE_NAMESPACE}
        )
    elseif (ALG_NAME STREQUAL "sphincs-a")
        target_compile_definitions(
            ${TARGET_NAME} PRIVATE
            SPHINCS_A_MODE=${MODE} 
            THASH=${ADD_ALGO_TEST_THASH}
            SPHINCS_A_HASH_MODE_NAMESPACE=${ADD_ALGO_TEST_MODE_NAMESPACE}
        )
    endif()
endfunction()

# Used for add bench test binary
function(add_algorithm_bench TARGET_NAME SOURCE_FILE ALG_NAME TYPE MODE)

    # Define other argument.
    set(options)
    set(oneValueArgs THASH MODE_NAMESPACE)
    set(multiValueArgs)

    # Parse argumanets
    cmake_parse_arguments(ADD_ALGO_TEST "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})

    if(NOT ADD_ALGO_TEST_THASH)
        set(ADD_ALGO_TEST_THASH "")
    endif()
    if(NOT ADD_ALGO_TEST_MODE_NAMESPACE)
        set(ADD_ALGO_TEST_MODE_NAMESPACE "")
    endif()

    add_executable(${TARGET_NAME} 
                        ${SOURCE_FILE} 
                        ${CMAKE_CURRENT_SOURCE_DIR}/utils/cpucycles.c
                        ${CMAKE_CURRENT_SOURCE_DIR}/utils/speed_print.c)
    target_link_libraries(${TARGET_NAME} PRIVATE pqmagic_static)  # bench use static lib
    add_dependencies(${TARGET_NAME} pqmagic_static_target)
    set_target_properties(${TARGET_NAME} PROPERTIES
        RUNTIME_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}"
    )
    # Include dir for using lib.
    target_include_directories(${TARGET_NAME} PRIVATE 
        ${CMAKE_CURRENT_SOURCE_DIR}/${TYPE}/${ALG_NAME}/${PQMAGIC_VERSION}
    )

    if (ALG_NAME STREQUAL "aigis-sig")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            AIGIS_SIG_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "dilithium")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            DILITHIUM_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "ml_dsa")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            ML_DSA_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "kyber")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            KYBER_K=${MODE}
        )
    elseif (ALG_NAME STREQUAL "aigis-enc")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            AIGIS_ENC_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "ml_kem")
        target_compile_definitions(${TARGET_NAME} PRIVATE 
            ML_KEM_MODE=${MODE}
        )
    elseif (ALG_NAME STREQUAL "slh_dsa")
        target_compile_definitions(
            ${TARGET_NAME} PRIVATE
            SLH_DSA_MODE=${MODE} 
            THASH=${ADD_ALGO_TEST_THASH}
            SLH_DSA_HASH_MODE_NAMESPACE=${ADD_ALGO_TEST_MODE_NAMESPACE}
        )
    elseif (ALG_NAME STREQUAL "sphincs-a")
        target_compile_definitions(
            ${TARGET_NAME} PRIVATE
            SPHINCS_A_MODE=${MODE} 
            THASH=${ADD_ALGO_TEST_THASH}
            SPHINCS_A_HASH_MODE_NAMESPACE=${ADD_ALGO_TEST_MODE_NAMESPACE}
        )
    endif()
endfunction()

# TEST and BENCH are default ON
option(ENABLE_TEST "Enable build test binary" ON)
option(ENABLE_BENCH "Enable build benchmark binary" ON)

if(ENABLE_TEST)
    message(STATUS "TEST BINARY BUILD: ENABLED")
else()
    message(STATUS "TEST BINARY BUILD: DISABLED")
endif()

if(ENABLE_BENCH)
    message(STATUS "BENCH BINARY BUILD: ENABLED")
else()
    message(STATUS "BENCH BINARY BUILD: DISABLED")
endif()

if(ENABLE_TEST OR ENABLE_BENCH)
    link_directories(${CMAKE_BINARY_DIR})
endif()

set(ALL_EXECUTABLE_TARGET "")
# KEM
if(ENABLE_KYBER)
    if(ENABLE_TEST)
        # Add Kyber test
        foreach(KYBER_MODE ${KYBER_MODES})
            add_algorithm_test(
                test_kyber_${KYBER_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/kem/kyber/${PQMAGIC_VERSION}/test/test_kyber.c"
                kyber kem ${KYBER_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET test_kyber_${KYBER_MODE})
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add Kyber bench
        foreach(KYBER_MODE ${KYBER_MODES})
            add_algorithm_bench(
                bench_kyber_${KYBER_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/kem/kyber/${PQMAGIC_VERSION}/test/test_speed.c" 
                kyber kem ${KYBER_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET bench_kyber_${KYBER_MODE})
        endforeach()
    endif()
endif()

if(ENABLE_ML_KEM)
    if(ENABLE_TEST)
        # Add Kyber test
        foreach(ML_KEM_MODE ${ML_KEM_MODES})
            add_algorithm_test(
                test_ml_kem_${ML_KEM_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/kem/ml_kem/${PQMAGIC_VERSION}/test/test_ml_kem.c"
                ml_kem kem ${ML_KEM_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET test_ml_kem_${ML_KEM_MODE})
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add Kyber bench
        foreach(ML_KEM_MODE ${ML_KEM_MODES})
            add_algorithm_bench(
                bench_ml_kem_${ML_KEM_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/kem/ml_kem/${PQMAGIC_VERSION}/test/test_speed.c" 
                ml_kem kem ${ML_KEM_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET bench_ml_kem_${ML_KEM_MODE})
        endforeach()
    endif()
endif()

if(ENABLE_AIGIS_ENC)
    if(ENABLE_TEST)
        # Add Aigis-enc test
        foreach(AIGIS_ENC_MODE ${AIGIS_ENC_MODES})
            add_algorithm_test(
                test_aigis_enc_${AIGIS_ENC_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/kem/aigis-enc/${PQMAGIC_VERSION}/test/test_aigis_enc.c"
                aigis-enc kem ${AIGIS_ENC_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET test_aigis_enc_${AIGIS_ENC_MODE})
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add Aigis-enc bench
        foreach(AIGIS_ENC_MODE ${AIGIS_ENC_MODES})
            add_algorithm_bench(
                bench_aigis_enc_${AIGIS_ENC_MODE} 
                "${CMAKE_CURRENT_SOURCE_DIR}/kem/aigis-enc/${PQMAGIC_VERSION}/test/test_speed.c" 
                aigis-enc kem ${AIGIS_ENC_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET bench_aigis_enc_${AIGIS_ENC_MODE})
        endforeach()
    endif()
endif()

# SIG
if(ENABLE_DILITHIUM)
    if(ENABLE_TEST)
        # Add Dilithium test
        foreach(DILITHIUM_MODE ${DILITHIUM_MODES})
            add_algorithm_test(
                test_dilithium_${DILITHIUM_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/sig/dilithium/${PQMAGIC_VERSION}/test/test_dilithium.c"
                dilithium sig ${DILITHIUM_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET test_dilithium_${DILITHIUM_MODE})
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add Dilithium bench
        foreach(DILITHIUM_MODE ${DILITHIUM_MODES})
            add_algorithm_bench(
                bench_dilithium_${DILITHIUM_MODE} 
                "${CMAKE_CURRENT_SOURCE_DIR}/sig/dilithium/${PQMAGIC_VERSION}/test/test_speed.c" 
                dilithium sig ${DILITHIUM_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET bench_dilithium_${DILITHIUM_MODE})
        endforeach()
    endif()
endif()

if(ENABLE_ML_DSA)
    if(ENABLE_TEST)
        # Add Dilithium test
        foreach(ML_DSA_MODE ${ML_DSA_MODES})
            add_algorithm_test(
                test_ml_dsa_${ML_DSA_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/sig/ml_dsa/${PQMAGIC_VERSION}/test/test_ml_dsa.c"
                ml_dsa sig ${ML_DSA_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET test_ml_dsa_${ML_DSA_MODE})
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add Dilithium bench
        foreach(ML_DSA_MODE ${ML_DSA_MODES})
            add_algorithm_bench(
                bench_ml_dsa_${ML_DSA_MODE} 
                "${CMAKE_CURRENT_SOURCE_DIR}/sig/ml_dsa/${PQMAGIC_VERSION}/test/test_speed.c" 
                ml_dsa sig ${ML_DSA_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET bench_ml_dsa_${ML_DSA_MODE})
        endforeach()
    endif()
endif()

if(ENABLE_AIGIS_SIG)
    if(ENABLE_TEST)
        # Add Aigis-sig test
        foreach(AIGIS_SIG_MODE ${AIGIS_SIG_MODES})
            add_algorithm_test(
                test_aigis_sig_${AIGIS_SIG_MODE}
                "${CMAKE_CURRENT_SOURCE_DIR}/sig/aigis-sig/${PQMAGIC_VERSION}/test/test_aigis.c"
                aigis-sig sig ${AIGIS_SIG_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET test_aigis_sig_${AIGIS_SIG_MODE})
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add Aigis-sig bench
        foreach(AIGIS_SIG_MODE ${AIGIS_SIG_MODES})
            add_algorithm_bench(
                bench_aigis_sig_${AIGIS_SIG_MODE} 
                "${CMAKE_CURRENT_SOURCE_DIR}/sig/aigis-sig/${PQMAGIC_VERSION}/test/test_speed.c" 
                aigis-sig sig ${AIGIS_SIG_MODE}
            )
            list(APPEND ALL_EXECUTABLE_TARGET bench_aigis_sig_${AIGIS_SIG_MODE})
        endforeach()
    endif()
endif()

if(ENABLE_SLH_DSA)
    if(ENABLE_TEST)
        # Add SLH-DSA test
        foreach(SLH_DSA_MODE ${SLH_DSA_MODES})
            set(MODE_NAMESPACE ${SLH_DSA_MODE})
            string(REPLACE "-" "_" MODE_NAMESPACE "${MODE_NAMESPACE}")
            string(REPLACE "slh_dsa_" "" MODE_NAMESPACE "${MODE_NAMESPACE}")

            if(PQMAGIC_VERSION STREQUAL "adv")
                message(FATAL_ERROR "Open Source Version Only Support PQMagic-std. Please set PQMAGIC_VERSION as \"std\" or contact as for further high performance support.")
            else()
                set(DIR_STR "${CMAKE_CURRENT_SOURCE_DIR}/sig/slh_dsa/${PQMAGIC_VERSION}/test/test_spx.c")
            endif()

            foreach(THASH ${SLH_DSA_THASH})
                add_algorithm_test(
                    test_slh_dsa_${MODE_NAMESPACE}_${THASH}
                    ${DIR_STR}
                    slh_dsa sig ${SLH_DSA_MODE} 
                    THASH ${THASH} MODE_NAMESPACE ${MODE_NAMESPACE}
                )
                list(APPEND ALL_EXECUTABLE_TARGET test_slh_dsa_${MODE_NAMESPACE}_${THASH})
            endforeach()
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add SLH-DSA bench
        foreach(SLH_DSA_MODE ${SLH_DSA_MODES})

            set(MODE_NAMESPACE ${SLH_DSA_MODE})
            string(REPLACE "-" "_" MODE_NAMESPACE "${MODE_NAMESPACE}")
            string(REPLACE "slh_dsa_" "" MODE_NAMESPACE "${MODE_NAMESPACE}")

            if(PQMAGIC_VERSION STREQUAL "adv")
                message(FATAL_ERROR "Open Source Version Only Support PQMagic-std. Please set PQMAGIC_VERSION as \"std\" or contact as for further high performance support.")
            else()
                set(DIR_STR "${CMAKE_CURRENT_SOURCE_DIR}/sig/slh_dsa/${PQMAGIC_VERSION}/test/test_speed.c")
            endif()

            foreach(THASH ${SLH_DSA_THASH})
                add_algorithm_bench(
                    bench_slh_dsa_${MODE_NAMESPACE}_${THASH}
                    ${DIR_STR}
                    slh_dsa sig ${SLH_DSA_MODE}
                    THASH ${THASH} MODE_NAMESPACE ${MODE_NAMESPACE}
                )
                list(APPEND ALL_EXECUTABLE_TARGET bench_slh_dsa_${MODE_NAMESPACE}_${THASH})
            endforeach()
        endforeach()
    endif()
endif()

if(ENABLE_SPHINCS_A)
    if(ENABLE_TEST)
        # Add SPHINCS-Alpha test
        foreach(SPHINCS_A_MODE ${SPHINCS_A_MODES})
            set(MODE_NAMESPACE ${SPHINCS_A_MODE})
            string(REPLACE "-" "_" MODE_NAMESPACE "${MODE_NAMESPACE}")
            string(REPLACE "sphincs_a_" "" MODE_NAMESPACE "${MODE_NAMESPACE}")

            if(PQMAGIC_VERSION STREQUAL "adv")
                message(FATAL_ERROR "Open Source Version Only Support PQMagic-std. Please set PQMAGIC_VERSION as \"std\" or contact as for further high performance support.")
            else()
                set(DIR_STR "${CMAKE_CURRENT_SOURCE_DIR}/sig/sphincs-a/${PQMAGIC_VERSION}/test/test_spx_a.c")
            endif()

            foreach(THASH ${SPHINCS_A_THASH})
                add_algorithm_test(
                    test_sphincs_a_${MODE_NAMESPACE}_${THASH}
                    ${DIR_STR}
                    sphincs-a sig ${SPHINCS_A_MODE} 
                    THASH ${THASH} MODE_NAMESPACE ${MODE_NAMESPACE}
                )
                list(APPEND ALL_EXECUTABLE_TARGET test_sphincs_a_${MODE_NAMESPACE}_${THASH})
            endforeach()
        endforeach()
    endif()

    if(ENABLE_BENCH)
        # Add SPHINCS-Alpha bench
        foreach(SPHINCS_A_MODE ${SPHINCS_A_MODES})

            set(MODE_NAMESPACE ${SPHINCS_A_MODE})
            string(REPLACE "-" "_" MODE_NAMESPACE "${MODE_NAMESPACE}")
            string(REPLACE "sphincs_a_" "" MODE_NAMESPACE "${MODE_NAMESPACE}")

            if(PQMAGIC_VERSION STREQUAL "adv")
                message(FATAL_ERROR "Open Source Version Only Support PQMagic-std. Please set PQMAGIC_VERSION as \"std\" or contact as for further high performance support.")
            else()
                set(DIR_STR "${CMAKE_CURRENT_SOURCE_DIR}/sig/sphincs-a/${PQMAGIC_VERSION}/test/test_speed.c")
            endif()

            foreach(THASH ${SPHINCS_A_THASH})
                add_algorithm_bench(
                    bench_sphincs_a_${MODE_NAMESPACE}_${THASH}
                    ${DIR_STR}
                    sphincs-a sig ${SPHINCS_A_MODE}
                    THASH ${THASH} MODE_NAMESPACE ${MODE_NAMESPACE}
                )
                list(APPEND ALL_EXECUTABLE_TARGET bench_sphincs_a_${MODE_NAMESPACE}_${THASH})
            endforeach()
        endforeach()
    endif()
endif()

# Custom install all executable target.
install(CODE 
    "
    if(NOT EXISTS ${INSTALL_BIN_DIR})
        file(MAKE_DIRECTORY ${INSTALL_BIN_DIR})
    endif()
    foreach(EXE_TARGET ${ALL_EXECUTABLE_TARGET})
       message(STATUS \"Installing: ${INSTALL_BIN_DIR}/\${EXE_TARGET}\")
       execute_process(COMMAND ln -s ${CMAKE_BINARY_DIR}/\${EXE_TARGET} ${INSTALL_BIN_DIR}/\${EXE_TARGET})
    endforeach()
    "
)

# Install pqmagic api
install(FILES include/pqmagic_api.h DESTINATION ${INSTALL_INCLUDE_DIR})

#########################
# Define clean all rule.#
#########################
add_custom_target(clean-all
    COMMAND ${CMAKE_COMMAND} -P ${CMAKE_BINARY_DIR}/cmake_clean_install.cmake
    COMMAND ${CMAKE_COMMAND} -P ${CMAKE_BINARY_DIR}/cmake_clean_target.cmake
    COMMENT "Cleaning build files and installed files..."
)

# Gen rules for removing installed stuff.
file(WRITE ${CMAKE_BINARY_DIR}/cmake_clean_install.cmake
    "if(EXISTS ${INSTALL_DIR})\n"
    "    message(STATUS \"Removing installed files from ${INSTALL_DIR}...\")\n"
    "    execute_process(COMMAND ${CMAKE_COMMAND} -E remove_directory ${INSTALL_DIR}/bin)\n"
    "    execute_process(COMMAND ${CMAKE_COMMAND} -E remove_directory ${INSTALL_DIR}/lib)\n"
    "    execute_process(COMMAND ${CMAKE_COMMAND} -E remove_directory ${INSTALL_DIR}/include)\n"
    "endif()\n"
    "message(STATUS \"Installed files removed.\")\n"
)

# Gen build target clean rules.
file(WRITE ${CMAKE_BINARY_DIR}/cmake_clean_target.cmake
    "message(STATUS \"Removing build files...\")\n"
    "execute_process(COMMAND ${CMAKE_COMMAND} --build ${CMAKE_BINARY_DIR} --target clean)\n"
    "file(GLOB_RECURSE SO_FILES \"./*.so\")\n"
    "file(GLOB_RECURSE A_FILES \"./*.a\")\n"
    "file(GLOB_RECURSE MRI_FILES \"./*.mri\")\n"
    "message(STATUS \"Removing shared library files...\")\n"
    "foreach(FILE \${SO_FILES})\n"
    "   file(REMOVE \${FILE})\n"
    "endforeach()\n"
    "message(STATUS \"Shared library files removed.\")\n"
    "message(STATUS \"Removing static library files...\")\n"
    "foreach(FILE \${A_FILES})\n"
    "   file(REMOVE \${FILE})\n"
    "endforeach()\n"
    "message(STATUS \"Static library files removed.\")\n"
    # "message(STATUS \"Removing .mri files...\")\n"
    # "foreach(FILE \${MRI_FILES})\n"
    # "   file(REMOVE \${FILE})\n"
    # "endforeach()\n"
    # "message(STATUS \".mri files removed.\")\n"
    "message(STATUS \"Build files removed.\")\n"
    "\n"
    "if(EXISTS ${INSTALL_BIN_DIR})\n"
    "    message(STATUS \"Removing files from ${INSTALL_BIN_DIR}...\")\n"
    "    file(GLOB BIN_FILES ${INSTALL_BIN_DIR}/*)\n"
    "    foreach(FILE \${BIN_FILES})\n"
    "        message(STATUS \"\tRemove \${FILE}.\")\n"
    "        file(REMOVE \${FILE})\n"
    "    endforeach()\n"
    "    message(STATUS \"Files removed from ${INSTALL_BIN_DIR}.\")\n"
    "endif()\n"
)

