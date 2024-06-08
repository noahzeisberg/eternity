<template>
  <Page>
    <div class="fixed bottom-0 left-0 bg-white dark:bg-black w-full flex justify-stretch p-3">
      <Dialog v-model:open="modalOpen">
        <DialogTrigger class="w-full">
          <Button class="w-full">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 mr-2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
            </svg>
            Hausaufgaben hinzufügen
          </Button>
        </DialogTrigger>
        <DialogContent class="">
          <DialogHeader>
            <DialogTitle>Neue Hausaufgabe</DialogTitle>
            <DialogDescription>Eine neue Hausaufgabe einstellen, die von jedem aufgerufen werden kann.</DialogDescription>
          </DialogHeader>
          <div class="flex flex-col gap-3">
            <Input v-model="subject" placeholder="Fach angeben..."></Input>

            <Textarea v-model="description" rows="5" placeholder="Hier Beschreibung angeben..."></Textarea>

            <Button @click="post()">Speichern</Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  </Page>
</template>
<script setup lang="ts">
const modalOpen = ref(false)

const subject = ref("")
const description = ref("")

function post() {
  modalOpen.value = false

  console.log({
    subject: subject.value,
    description: description.value,
    date: new Date().toISOString()
  })
}
</script>