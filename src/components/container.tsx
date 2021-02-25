import React from "react"
import { Box, BoxProps } from "@chakra-ui/react"

export const Container = (props: BoxProps) => (
  <Box
    w="full"
    pb="12"
    pt="3"
    maxW="1200px"
    mx="auto"
    px={[0,0,4,8]}
    {...props}
  />
)

export default Container
