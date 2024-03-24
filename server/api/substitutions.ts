export default defineEventHandler(async (event) => {
    const query = getQuery(event)
    const preferredClass = query.class?.toString()
    if(preferredClass === null || preferredClass === undefined) {
        return { msg: "Class not given." }
    }

    const payloadWrapper = await getPayload()
    const payload = payloadWrapper.payload

    let rows: any[] = []
    payload.rows.forEach(item => {
        rows.push({
            "hour": item.data[0],
            "class": item.data[1],
            "subject": item.data[2],
            "room": sanitize(item.data[3]).split(" ")[0],
            "teacher": sanitize(item.data[4]),
            "type": item.data[5]
        })
    })

    return rows.filter((substitution) => substitution["class"].toLowerCase().includes(preferredClass.toLowerCase()))
})

const sanitize = (string: string): string => string.replace(/<[^>]*>?/gm, "");

const getPayload = async () => await $fetch<{
    payload: { rows: { data: any }[], date: string, lastUpdate: string, weekDay: string }
}>("https://ajax.webuntis.com/WebUntis/monitor/substitution/data?school=hbs-hattersheim", {
    method: "POST",
    body: {
        "formatName": "Vetr_Kla",
        "schoolName": "hbs-hattersheim",
        "date": formatDate(new Date()),
        "dateOffset": 0,
        "strikethrough": false,
        "mergeBlocks": false,
        "showOnlyFutureSub": false,
        "showBreakSupervisions": false,
        "showTeacher": true,
        "showClass": true,
        "showHour": true,
        "showInfo": true,
        "showRoom": true,
        "showSubject": true,
        "groupBy": 1,
        "hideAbsent": false,
        "departmentIds": [],
        "departmentElementType": -1,
        "hideCancelWithSubstitution": true,
        "hideCancelCausedByEvent": true,
        "showTime": false,
        "showSubstText": false,
        "showAbsentElements": [],
        "showAffectedElements": [1],
        "showUnitTime": true,
        "showMessages": true,
        "showStudentgroup": false,
        "enableSubstitutionFrom": false,
        "showSubstitutionFrom": 1800,
        "showTeacherOnEvent": false,
        "showAbsentTeacher": false,
        "strikethroughAbsentTeacher": false,
        "activityTypeIds": [2, 3, 4, 8],
        "showEvent": true,
        "showCancel": true,
        "showOnlyCancel": false,
        "showSubstTypeColor": false,
        "showExamSupervision": false,
        "showUnheraldedExams": false
    }
});

const formatDate = (inputDate: Date) => inputDate.getFullYear().toString() + pad(inputDate.getMonth() + 1, 2).toString() + pad(inputDate.getDate(), 2).toString();

const pad = (number: number, length: number) => {
    let str = '' + number;
    while (str.length < length) {
        str = '0' + str;
    }
    return str;
};