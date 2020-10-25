add_executable("HuffManCompression" huffmanComp.cpp)
add_custom_command(TARGET HuffManCompression POST_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${PROJECT_SOURCE_DIR}/compression/huffman/text.txt"
        $<TARGET_FILE_DIR:HuffManCompression>)
