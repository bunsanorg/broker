cmake_minimum_required(VERSION 3.0)

include(${CMAKE_SOURCE_DIR}/system-config.cmake OPTIONAL)

project(bunsan_broker)

if(BunsanCMakeBuiltin)
    BunsanCMakeLoad()
else()
    find_package(BunsanCMake CONFIG REQUIRED)
endif()

bunsan_find_python3()

bunsan_add_protobuf_cxx_library(
    TARGET ${PROJECT_NAME}_pb
    PYTHON
    PROTOS
        bunsan/broker/connection.proto
        bunsan/broker/protocol.proto
        bunsan/broker/rabbit/connection.proto
    INSTALL
)
bunsan_install_python_package(
    PACKAGE bunsan.broker
    DIRECTORY pysrc
)
bunsan_add_library(${PROJECT_NAME}
    src/task/stream_channel.cpp
)
bunsan_use_target(${PROJECT_NAME} ${PROJECT_NAME}_pb)
bunsan_use_bunsan_package(${PROJECT_NAME} bunsan_crypto bunsan_crypto)
bunsan_use_bunsan_package(${PROJECT_NAME} bunsan_protobuf bunsan_protobuf)

bunsan_install_headers()
bunsan_install_targets(
    ${PROJECT_NAME}
)
bunsan_install_project()

bunsan_include_tests()
