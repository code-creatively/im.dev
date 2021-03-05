import React, { useState } from "react"
import { Avatar, Divider, Flex, Heading, HStack, IconButton, Text, VStack, chakra, Menu, MenuButton, MenuList, MenuItem, Box } from "@chakra-ui/react"
import { Comment } from "src/types/comments"
import Card from "components/card"
import { getUserName } from "utils/user"
import moment from 'moment'
import { MarkdownRender } from "components/markdown-editor/render"
import Like from "components/interaction/like"
import { FaRegEdit, FaRegFlag, FaRegTrashAlt, FaReply, FaTrash } from "react-icons/fa"
import { User } from "src/types/user"
import CommentEditor from "./editor"
import { requestApi } from "utils/axios/request"
import Link from "next/link"

interface Props {
    user: User
    comment: Comment
    parent: Comment
    onChange: any
}
export const Reply = (props: Props) => {
    const { comment, user,onChange,parent} = props
    const [editorVisible, setEditorVisible] = useState(false)

    const [replyVisible,setReplyVisible] = useState(false)
    const [reply,setReply] = useState('')
    const submitReply = async (md) => {
        await requestApi.post('/story/comment',{targetID: parent.id, md: md})
        setReplyVisible(false);
        onChange()
    }

    const changeReply = async (md) => {
        await requestApi.post('/story/comment',{...comment, md})
        setEditorVisible(false)
        onChange()
    }

    const deleteReply = async id => {
        await requestApi.delete(`/story/comment/${id}`)
        onChange()
    }

    const likeReply = async (id) => {
        await requestApi.post(`/story/like/${id}`)
        onChange()
      }


    const replyToReply = () => {
        if (comment.creator.nickname === "") {
            setReply(`@${comment.creator.username}`)
        } else {
            setReply(`@[${comment.creator.nickname}](/${comment.creator.username})`)
        }
        
        setReplyVisible(!replyVisible)
    }

    return (
        <>{
            editorVisible ? (user && <CommentEditor user={user} md={comment.md} onSubmit={md => {setEditorVisible(false);changeReply(md)}} onCancel={() => setEditorVisible(false)} menu={false} />) :
                    <VStack alignItems="left">
                        <Flex width="100%" alignItems="center" justifyContent="space-between">
                            <Link href={`/${comment.creator.username}`}>
                                <HStack spacing="4" cursor="pointer">
                                    <Avatar src={comment.creator.avatar} width="40px" height="40px"></Avatar>
                                    <VStack alignItems="left">
                                        <Heading size="sm">{getUserName(comment.creator)}</Heading>
                                    </VStack>
                                </HStack>
                            </Link>
                            <Text layerStyle="textSecondary" ml="2" fontSize=".9rem">{moment(comment.created).fromNow()}</Text>
                        </Flex>

                        <MarkdownRender md={comment.md} pl="16" pr="2" mt="3" />

                        <Flex justifyContent="space-between" pl="16" pr="2">
                            <Like liked={comment.liked} count={comment.likes} storyID={comment.id} />
                            <HStack>
                                {user && <IconButton
                                    aria-label="go to github"
                                    variant="ghost"
                                    _focus={null}
                                    color="gray.500"
                                    icon={<FaReply />}
                                    onClick={replyToReply}
                                    fontSize="18px"
                                />}

                                <Menu>
                                    <IconButton
                                        as={MenuButton}
                                        aria-label="go to github"
                                        variant="ghost"
                                        _focus={null}
                                        color="gray.500"
                                        icon={<chakra.svg fill='gray.500' height="21px" viewBox="0 0 512 512"><path d="M256 8C119 8 8 119 8 256s111 248 248 248 248-111 248-248S393 8 256 8zm216 248c0 118.7-96.1 216-216 216-118.7 0-216-96.1-216-216 0-118.7 96.1-216 216-216 118.7 0 216 96.1 216 216zm-207.5 86.6l115-115.1c4.7-4.7 4.7-12.3 0-17l-7.1-7.1c-4.7-4.7-12.3-4.7-17 0L256 303l-99.5-99.5c-4.7-4.7-12.3-4.7-17 0l-7.1 7.1c-4.7 4.7-4.7 12.3 0 17l115 115.1c4.8 4.6 12.4 4.6 17.1-.1z"></path></chakra.svg>}
                                        _hover={null}
                                        _active={null}
                                    />
                                    <MenuList>
                                        {user && <MenuItem onClick={() => setEditorVisible(true)}><FaRegEdit /><chakra.span ml="2">Edit</chakra.span></MenuItem>}
                                        {user && <MenuItem onClick={() => deleteReply(comment.id)}><FaRegTrashAlt /><chakra.span ml="2">Delete</chakra.span></MenuItem>}
                                        <MenuItem><FaRegFlag /><chakra.span ml="2">Report</chakra.span></MenuItem>
                                    </MenuList>
                                </Menu>
                            </HStack>
                        </Flex>

                        {replyVisible && 
                            <Box pl="16" pr="2">
                                <CommentEditor user={user} md={reply} onSubmit={md => {submitReply(md)}} onCancel={() => setReplyVisible(false)} menu={false}/>
                            </Box>}
                    </VStack>}
        </>
    )
}

export default Reply
