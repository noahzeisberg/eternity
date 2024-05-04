<template>
  <Page>
    <div class="fixed bottom-0 left-0 bg-white dark:bg-black w-full flex justify-stretch p-3">
      <Dialog>
        <DialogTrigger class="w-full">
          <Button class="w-full">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 mr-1 invert">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
            </svg>
            Hausaufgaben hinzufügen
          </Button>
        </DialogTrigger>
        <DialogContent class="">
          <DialogHeader>
            <DialogTitle>Neue Hausaufgabe</DialogTitle>
            <DialogDescription>Eine neue Hausaufgabe einstellen, die von Schülern aufgerufen werden kann.</DialogDescription>
          </DialogHeader>
          <div class="flex flex-col gap-3">
            <Select>
              <SelectTrigger>
                <SelectValue placeholder="Fach auswählen..." />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="apple">Buch</SelectItem>
              </SelectContent>
            </Select>

            <Select>
              <SelectTrigger>
                <SelectValue placeholder="Typ auswählen..." />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="apple">Buch</SelectItem>
              </SelectContent>
            </Select>

            <Textarea v-model="description" rows="5" placeholder="Hier Beschreibung angeben."></Textarea>

            <Button>Speichern</Button>
            <Separator></Separator>
            <CameraInput v-model="image">Foto hochladen</CameraInput>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  </Page>
</template>
<script setup lang="ts">
import { createWorker } from 'tesseract.js';
import { toast } from "vue-sonner";

const image = ref()

watch(image, () => {
  console.log(image.value)
  toast.promise(extractText(image.value), {
    loading: "Wird importiert...",
    success: (data) => {
      description.value = data
      return "Beschreibung importiert!"
    }
  })
})

const description = ref()

async function extractText(image: File | string): Promise<string> {
  const worker = await createWorker("eng+deu")
  const output = await worker.recognize(image)
  await worker.terminate()
  return output.data.text
}
</script>