import { requestApi } from "../src/utils/axios/request"

export let config = {
    appName: "im.dev",
    uiDomain: "http://localhost:4004",
    commonMaxlen: 255,
    posts: {
        titleMaxLen: 128,
        briefMaxLen: 128,
        writingEnabled: false,
        maxTags: 0
    },
    user: {
        nicknameMaxLen: 64,
        usernameMaxLen: 39,
        navbarMaxLen: 20
    }
}

export function initUIConfig() {
    requestApi.get("/config").then((res) => {
        console.log("初始化UI config:", res.data)
        config = res.data
    })}