export default defineEventHandler(async (): Promise<string[]> => {
    let cycle
    const getCurrentWeek = () => {
        const now: Date = new Date()
        const startOfYear: Date = new Date(now.getFullYear(), 0, 0)
        const diff: number = now.getTime() - startOfYear.getTime()
        const oneWeekInMillis: number = 1000 * 60 * 60 * 24 * 7
        return Math.floor(diff / oneWeekInMillis)
    }

    const students: string[] = await $fetch("/api/students")
    let groups: string[][] = []
    if(students.length % 2 !== 0) {
        students.push(...students)
    }

    let even: string[] = []
    let odd: string[] = []

    for (let i: number = 0; i < students.length; i++) {
        if(i % 2 === 0) {
            even.push(students[i])
        } else {
            odd.push(students[i])
        }
    }

    for (let i: number = 0; i < even.length; i++) {
        groups.push([even[i], odd[i]])
    }

    if (getCurrentWeek() >= groups.length) {
        cycle = getCurrentWeek() % groups.length
    } else {
        cycle = getCurrentWeek()
    }

    return groups[cycle]
})