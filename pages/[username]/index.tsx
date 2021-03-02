import { Box, chakra, Flex, HStack, VStack, Image, Heading, Text, Button, useColorModeValue, Divider } from "@chakra-ui/react"
import Card from "components/card"
import Container from "components/container"
import SEO from "components/seo"
import siteConfig from "configs/site-config"
import useSession from "hooks/use-session"
import Nav from "layouts/nav/nav"
import VerticalNav from "layouts/nav/vertical-nav"
import PageContainer from "layouts/page-container"
import PageContainer1 from "layouts/page-container1"
import { useRouter } from "next/router"
import React, { useEffect, useState } from "react"
import { FaDove, FaEdit, FaFacebook, FaGithub, FaHeart, FaPlus, FaRegStar, FaStackOverflow, FaStar, FaTwitter, FaWeibo, FaZhihu } from "react-icons/fa"
import { ReserveUrls } from "src/data/reserve-urls"
import { User } from "src/types/session"
import { requestApi } from "utils/axios/request"
import moment from 'moment'

const UserPage = () => {
  const router = useRouter()
  const username = router.query.username
  const session = useSession()
  const [user, setUser]: [User, any] = useState(null)
  const borderColor = useColorModeValue('white', 'transparent')
  useEffect(() => {
    if (username) {
      requestApi.get(`/user/info/${username}`).then(res => setUser(res.data))
    }
  }, [username])
  return (
    <>
      <SEO
        title={siteConfig.seo.title}
        description={siteConfig.seo.description}
      />
      <PageContainer1>
        {
          user &&
          <Box alignItems="left">
            <Card p="0" borderTop="none">
              <Box backgroundImage={`url(${user.cover})`} height="300px" width="100%" backgroundSize="cover" backgroundPosition="center" />
              <VStack maxHeight="200px" position="relative" top="-70px" spacing="3">
                <Image src={user.avatar} height="130px" borderRadius="50%" border={`4px solid ${borderColor}`} />
                <Heading fontSize="1.8rem">{user.nickname}</Heading>
                {user.tagline && <Text layerStyle="textSecondary" fontWeight="450" fontSize="1.2rem" ml="1" mt="2">{user.tagline}</Text>}
                <Flex layerStyle="textSecondary" spacing="2" pt="3" alignItems="center">
                  <chakra.span><FaHeart /></chakra.span><chakra.span ml="1">Followers <chakra.a fontWeight="600">0</chakra.a></chakra.span>
                  <chakra.span ml="5"><FaStar /></chakra.span><chakra.span ml="1">Following <chakra.a fontWeight="600">0</chakra.a></chakra.span>
                </Flex>
                <Box pt="3" position="absolute" right="15px" top="60px">{session?.user.id === user.id ? <Button onClick={() => router.push(`${ReserveUrls.Settings}/profile`)} variant="outline" leftIcon={<svg height="1.3rem" fill="currentColor" viewBox="0 0 512 512"><path d="M493.255 56.236l-37.49-37.49c-24.993-24.993-65.515-24.994-90.51 0L12.838 371.162.151 485.346c-1.698 15.286 11.22 28.203 26.504 26.504l114.184-12.687 352.417-352.417c24.992-24.994 24.992-65.517-.001-90.51zm-95.196 140.45L174 420.745V386h-48v-48H91.255l224.059-224.059 82.745 82.745zM126.147 468.598l-58.995 6.555-30.305-30.305 6.555-58.995L63.255 366H98v48h48v34.745l-19.853 19.853zm344.48-344.48l-49.941 49.941-82.745-82.745 49.941-49.941c12.505-12.505 32.748-12.507 45.255 0l37.49 37.49c12.506 12.506 12.507 32.747 0 45.255z"></path></svg>}>Edit Profile</Button>
                  : <Button colorScheme="teal">Follow</Button>}</Box>
              </VStack>
            </Card>
            <HStack spacing="4" mt="4">
              <VStack alignItems="left" spacing="4" width="350px" display={{ base: "none", md: "flex" }}>
                <Card>
                  <Text layerStyle="textSecondary">{user.about}</Text>
                  <Divider my="4" />
                  {user.location && <HStack>
                    <chakra.span layerStyle="textSecondary" width="90px">Location: </chakra.span>
                    <chakra.span fontWeight="500" ml="2">{user.location}</chakra.span>
                  </HStack>}

                  {user.website && <HStack mt="1">
                    <chakra.span layerStyle="textSecondary" width="90px">Website: </chakra.span>
                    <chakra.a fontWeight="500" ml="2" href={user.website} target="_blank">{user.website}</chakra.a>
                  </HStack>}

                  <HStack mt="1">
                    <chakra.span layerStyle="textSecondary" width="90px">Joined: </chakra.span>
                    <chakra.span fontWeight="500" ml="2">{moment(user.created).fromNow()}</chakra.span>
                  </HStack>

                  <HStack layerStyle="textSecondary" fontSize="1.4rem" mt="4" spacing="5">
                    {user.github && <chakra.a href={user.github} target="_blank"><FaGithub /></chakra.a>}
                    {user.twitter && <chakra.a href={user.twitter} target="_blank"><FaTwitter /></chakra.a>}
                    {user.facebook && <chakra.a href={user.facebook} target="_blank"><FaFacebook /></chakra.a>}
                    {user.stackoverflow && <chakra.a href={user.stackoverflow} target="_blank"><FaStackOverflow /></chakra.a>}
                    {user.zhihu && <chakra.a href={user.zhihu} target="_blank"><FaZhihu /></chakra.a>}
                    {user.weibo && <chakra.a href={user.weibo} target="_blank"><FaWeibo /></chakra.a>}
                  </HStack>
                  <Divider my="4" />

                  {user.availFor && <Box>
                    <Text fontWeight="600" layerStyle="textSecondary">I am available for</Text>
                    <Text mt="2">{user.availFor}</Text>
                  </Box>}
                </Card>
              </VStack>
            </HStack>
          </Box>
        }
      </PageContainer1>
    </>
  )
}

export default UserPage

