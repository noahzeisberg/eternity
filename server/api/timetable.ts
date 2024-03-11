export default defineEventHandler(async () => {
    return {
        timetable: [
            {
                day: "Montag",
                lessons: [
                    { name: "Physik", teacher: "STN", room: 129 },
                    { name: "Physik", teacher: "STN", room: 129 },
                    { name: "Englisch", teacher: "FTH", room: 221 },
                    { name: "Englisch", teacher: "FTH", room: 221 },
                    { name: "Deutsch", teacher: "PLC", room: 221 },
                    { name: "Deutsch", teacher: "PLC", room: 221 },
                ]
            },
            {
                day: "Dienstag",
                lessons: [
                    { name: "Kunst", teacher: "RSH", room: 0 },
                    { name: "Kunst", teacher: "RSH", room: 0 },
                    { name: "Chemie", teacher: "PTT", room: 0 },
                    { name: "Chemie", teacher: "PTT", room: 0 },
                    { name: "Mathe", teacher: "NCH", room: 221 },
                    { name: "Mathe", teacher: "NCH", room: 221 },
                ]
            },
            {
                day: "Mittwoch",
                lessons: [
                    { name: "Mathe", teacher: "NCH", room: 221 },
                    { name: "Mathe", teacher: "NCH", room: 221 },
                    { name: "Sport", teacher: "NHST", room: "SPO" },
                    { name: "Sport", teacher: "NHST", room: "SPO" },
                    { name: "Spanisch", teacher: "DRR", room: 207 },
                    { name: "Spanisch", teacher: "DRR", room: 207 },
                ]
            },
            {
                day: "Donnerstag",
                lessons: [
                    { name: "PoWi", teacher: "LPF", room: 221 },
                    { name: "PoWi", teacher: "LPF", room: 221 },
                    { name: "Religion", teacher: "PLC", room: 218 },
                    { name: "Religion", teacher: "PLC", room: 218 },
                    { name: "Deutsch", teacher: "PLC", room: 221 },
                    { name: "Deutsch", teacher: "PLC", room: 221 },
                ]
            },
            {
                day: "Freitag",
                lessons: [
                    { name: "Geschichte", teacher: "???", room: 221 },
                    { name: "Geschichte", teacher: "???", room: 221 },
                    { name: "Englisch", teacher: "FTH", room: 221 },
                    { name: "Englisch", teacher: "FTH", room: 221 },
                    { name: "Spanisch", teacher: "DRR", room: 207 },
                    { name: "Spanisch", teacher: "DRR", room: 207 },
                ]
            },
        ]
    }
})