export default defineEventHandler(async event => {
    assertMethod(event, "POST")
    const storage = useStorage("dataStorage")
    const body = event._requestBody?.toString()

})