if(EXISTS /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs)
    message(STATUS "Removing installed files from /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs...")
    execute_process(COMMAND /usr/bin/cmake -E remove_directory /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/bin)
    execute_process(COMMAND /usr/bin/cmake -E remove_directory /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/lib)
    execute_process(COMMAND /usr/bin/cmake -E remove_directory /home/teddycode/Desktop/Workspace/crypto-suites/bin/webassembly/cmd/wasm/pqcgo/libs/include)
endif()
message(STATUS "Installed files removed.")
