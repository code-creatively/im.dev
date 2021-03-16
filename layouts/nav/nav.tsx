import {
  chakra,
  Flex,
  HStack,
  IconButton,
  useColorModeValue,
  useDisclosure,
  useUpdateEffect,
  Box
} from "@chakra-ui/react"
import siteConfig from "configs/site-config"
import { useViewportScroll } from "framer-motion"
import NextLink from "next/link"
import React, { useEffect, useState } from "react"
import { FaGithub } from "react-icons/fa"
import Logo, { LogoIcon } from "src/components/logo"
import { MobileNavButton, MobileNavContent } from "./mobile-nav"
import AlgoliaSearch from "src/components/search/algolia-search"
import { useRouter } from "next/router"
import { ReserveUrls } from "src/data/reserve-urls"
import Link from "next/link"
import DarkMode from "components/dark-mode"
import AccountMenu from "components/user-menu"
import { navLinks } from "src/data/links"
import { getSvgIcon } from "components/svg-icon"
import { requestApi } from "utils/axios/request"


function HeaderContent() {
  const router = useRouter()
  const { asPath } = router

  const mobileNav = useDisclosure()


  const mobileNavBtnRef = React.useRef<HTMLButtonElement>()

  const [navs,setNavs] = useState(navLinks)
  useEffect(() => {
    requestApi.get("/navbars").then(res => {
       const nvs = []
       res.data.forEach(nv => nvs.push({
         title: nv.label,
         url: nv.value
       }))

       setNavs(nvs)
    })
  },[])

  useUpdateEffect(() => {
    mobileNavBtnRef.current?.focus()
  }, [mobileNav.isOpen])


  return (
    <>
      <Flex w="100%" h="100%" align="center" justify="space-between" px={{ base: "4", md: "6" }}>
        <Flex align="center">
          <NextLink href="/" passHref>
            <chakra.a display={{ base: "none", md: "block" }} style={{ marginTop: '-5px' }} aria-label="Chakra UI, Back to homepage">
              <Logo width="130" />
            </chakra.a>
          </NextLink>
          <NextLink href="/" passHref>
            <chakra.a display={{ base: "block", md: "none" }} aria-label="Chakra UI, Back to homepage">
              <LogoIcon />
            </chakra.a>
          </NextLink>

          <HStack ml={{ base: 1, md: 4, lg: 12 }} fontSize="1rem">
            {navs.map(link => <Box px={[0,0,4,4]} py="0.7rem" rounded="md" key={link.url} color={useColorModeValue("gray.700", "whiteAlpha.900")} aria-current={asPath === link.url ? "page" : undefined} _activeLink={{ bg: useColorModeValue("transparent", "rgba(48, 140, 122, 0.3)"), color: useColorModeValue("teal.500", "teal.200"), fontWeight: "bold", }} ><Link href={link.url}>{link.title}</Link></Box>)}
          </HStack>
        </Flex>

        <HStack
          // w="100%"
          maxW="600px"
          align="center"
          color={useColorModeValue("gray.500", "gray.400")}
          spacing="1"
        >
          {/* <AlgoliaSearch />  */}
          <Link
            aria-label="Go to Chakra UI GitHub page"
            href={siteConfig.repo.url}
          >
            <IconButton
              size="md"
              fontSize="lg"
              aria-label="go to github"
              variant="ghost"
              color="current"
              _focus={null}
              display={{ base: "none", md: "block" }}
              icon={<FaGithub />}
            />
          </Link>
          <Link
            aria-label="Go to Chakra UI GitHub page"
            href={`${ReserveUrls.Search}/posts`}
          >
            <IconButton
              aria-label="go to github"
              variant="ghost"
              color="current"
              _focus={null}
              icon={getSvgIcon("search")}
            />
          </Link>
          <DarkMode />
          <AccountMenu />
          {/* <MobileNavButton
            ref={mobileNavBtnRef}
            aria-label="Open Menu"
            onClick={mobileNav.onOpen}
          /> */}
        </HStack>
      </Flex>
      {/* <MobileNavContent isOpen={mobileNav.isOpen} onClose={mobileNav.onClose} /> */}
    </>
  )
}

function Header(props) {
  const ref = React.useRef<HTMLHeadingElement>()
  const [y, setY] = React.useState(0)
  const { height = 0 } = ref.current?.getBoundingClientRect() ?? {}

  const { scrollY } = useViewportScroll()
  React.useEffect(() => {
    return scrollY.onChange(() => setY(scrollY.get()))
  }, [scrollY])

  return (
    <chakra.header
      ref={ref}
      shadow={y > height ? "sm" : undefined}
      transition="box-shadow 0.2s"
      pos="fixed"
      top="0"
      zIndex="3"
      left="0"
      right="0"
      width="full"
      bg={useColorModeValue('gray.50', 'gray.800')}
      {...props}
    >
      <chakra.div height="4.5rem" mx="auto" maxW="1200px">
        <HeaderContent />
      </chakra.div>
    </chakra.header>
  )
}

export default Header

