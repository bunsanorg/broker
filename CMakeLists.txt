cmake_minimum_required(VERSION 3.0)

include(${CMAKE_SOURCE_DIR}/system-config.cmake OPTIONAL)

project(bunsan_broker)

find_package(BunsanCMake CONFIG REQUIRED)

bunsan_add_protobuf_cxx_library(
    TARGET ${PROJECT_NAME}_pb
    PROTOS
        bunsan/broker/client.proto
    INSTALL
)
#bunsan_add_library(${PROJECT_NAME}
#)
#bunsan_use_target(${PROJECT_NAME} ${PROJECT_NAME}_pb)
#bunsan_use_bunsan_package(${PROJECT_NAME} bunsan_protobuf bunsan_protobuf)

#bunsan_add_executable(${PROJECT_NAME}_cli
#    src/bin/cli.cpp
#)
#bunsan_use_target(${PROJECT_NAME}_cli ${PROJECT_NAME})

bunsan_install_headers()
bunsan_install_targets(
    #${PROJECT_NAME}
    #${PROJECT_NAME}_cli
)
bunsan_install_project()

bunsan_include_tests()