import React from "react"
import { 
    IconButton,  
    Menu,
    MenuButton,
    MenuList,
    MenuItem,
    MenuDivider,
    Image,
    Button
} from "@chakra-ui/react"
import useSession from "hooks/use-session"
import { Session } from "src/types/user"
import { useRouter } from "next/router"
import storage from "utils/localStorage"
import { ReserveUrls } from "src/data/reserve-urls"
import { FaRegSun, FaUserAlt ,FaBookmark, FaSignOutAlt,FaEdit,FaStar, FaHeart} from "react-icons/fa"
import { isAdmin, isEditor } from "utils/role"
import { logout } from "utils/session"
import Link from "next/link"

export const UserMenu = () => {
    const session: Session = useSession()
    const router = useRouter()

    const login = () => {
      console.log(router)
      storage.set("current-page", router.asPath)
      router.push(ReserveUrls.Login)
    }

    return (
        <>
            {session ?
              <Menu>
                <MenuButton
                  as={IconButton}
                  bg="transparent"
                  _focus={null}
                  icon={session.user.avatar !== '' ? <Image
                    boxSize="2.8em"
                    borderRadius="full"
                    src={session.user.avatar}
                    alt="user"
                  /> :
                    <FaUserAlt />
                  }
                  aria-label="Options"
                  ml={{ base: "0", md: "2" }}
                />
                <MenuList>
                  <Link href={`/${session.user.username}`}>
                    <MenuItem icon={<FaUserAlt fontSize="16" />}>
                      <span>{session.user.nickname}</span>
                    </MenuItem>
                  </Link>
                  <MenuDivider />
                  {isEditor(session.user.role) && <Link href={`${ReserveUrls.Editor}/posts`}><MenuItem icon={<FaEdit fontSize="16" />} >创作中心</MenuItem></Link>}
                  <Link href={`${ReserveUrls.Bookmarks}`}><MenuItem icon={<FaBookmark fontSize="16" />}>书签收藏</MenuItem></Link>
                  <Link href={`${ReserveUrls.Interaction}/following-tags`}><MenuItem icon={<FaHeart fontSize="16" />}>我的关注</MenuItem></Link>
                  {isAdmin(session.user.role) && <Link href={`${ReserveUrls.Admin}/tags`}><MenuItem  icon={<FaStar fontSize="16" />} >管理员</MenuItem></Link>}
                  <MenuDivider />
                  <Link href={`${ReserveUrls.Settings}/profile`}><MenuItem icon={<FaRegSun fontSize="16" />}>偏好设置</MenuItem></Link>
                  <MenuItem onClick={() => logout()} icon={<FaSignOutAlt fontSize="16" />}>账号登出</MenuItem>
                </MenuList>
              </Menu> :
              <Button
                as="a"
                ml="2"
                colorScheme="teal"
                fontSize=".8rem"
                onClick={() => login()}
              >
                SIGN IN
              </Button>
            }
        </>
    )
}

export default UserMenu
