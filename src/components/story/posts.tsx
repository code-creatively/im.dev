import React from "react"
import { Box, Center, Text, useColorModeValue, VStack } from "@chakra-ui/react"
import { Post } from "src/types/posts"
import PostCard from "./post-card"
import userCustomTheme from "theme/user-custom"

interface Props {
    posts: Post[]
    card?: any
    size?: 'sm' | 'md'
    showFooter?: boolean
    type?: string
    highlight?: string
}


export const Posts = (props: Props) => {
    const { posts,card=PostCard,showFooter=true,type="classic"} = props
    const postBorderColor = useColorModeValue(userCustomTheme.borderColor.light, userCustomTheme.borderColor.dark)
    const Card = card
    const showBorder = i => {
        if (i < posts.length -1) {
            return true
        }

        if (showFooter) {
            return true 
        } else {
            return false
        }
    }
    return (
        <>
            <VStack alignItems="left">
                {posts.map((post,i) =>
                    <Box py="2" borderBottom={showBorder(i)? `1px solid ${postBorderColor}`:null} key={post.id}>
                        <Card post={post} size={props.size} type={type} highlight={props.highlight}/>
                    </Box>)}
            </VStack>
            {showFooter && <Center><Text layerStyle="textSecondary" fontSize="sm" py="4">没有更多文章了</Text></Center>}
        </>
    )
}

export default Posts
