<template>
  <div>
    <div class="flex justify-between items-center w-full p-3">
      <Sheet>
        <SheetTrigger>
          <Button variant="ghost" size="icon">
            <HamburgerMenuIcon class="w-4 h-4"/>
          </Button>
        </SheetTrigger>
        <SheetContent side="left">
          <SheetHeader>
            <SheetTitle>Are you sure absolutely sure?</SheetTitle>
            <SheetDescription>
              This action cannot be undone. This will permanently delete your account
              and remove your data from our servers.
            </SheetDescription>
          </SheetHeader>
          <SheetFooter></SheetFooter>
        </SheetContent>
      </Sheet>

      <h1 class="font-black">ETERNITY</h1>

      <Button variant="ghost" size="icon">
        <GearIcon class="w-4 h-4"/>
      </Button>
    </div>
    <NuxtPage></NuxtPage>
  </div>
</template>

<script setup lang="ts">
import { HamburgerMenuIcon, GearIcon } from '@radix-icons/vue'
</script>