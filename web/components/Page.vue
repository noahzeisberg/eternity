<template>
  <div>
    <nav class="bg-white dark:bg-black border-b border-zinc-100 dark:border-zinc-800 flex justify-between items-center w-full p-3">
      <div class="flex items-center gap-2">
        <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="size-8 mx-1.5">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M11 8L3 16H11L11 8ZM13 16L21 8L13 8V16Z" fill="currentColor"/>
        </svg>
        <h1 class="text-xs font-black">{{ pageName === "" ? "ETERNITY" : pageName }}</h1>
      </div>

      <Sheet>
        <SheetTrigger>
          <Button variant="ghost" size="icon">
            <HamburgerMenuIcon/>
          </Button>
        </SheetTrigger>
        <SheetContent class="w-80 h-full">
          <SheetHeader>
            <SheetTitle>Willkommen!</SheetTitle>
            <SheetDescription>Hier ist eine Übersicht von den verschiedenen Bereichen.</SheetDescription>
          </SheetHeader>
          <Separator class="my-3"></Separator>
          <div class="flex flex-col gap-0.5">
            <SidebarItem name="Übersicht" to="/">
              <HomeIcon></HomeIcon>
            </SidebarItem>

            <SidebarItem name="Vertretungen" to="/substitutions">
              <CircleBackslashIcon></CircleBackslashIcon>
            </SidebarItem>

            <SidebarItem name="Kalender" to="/calendar">
              <CalendarIcon></CalendarIcon>
            </SidebarItem>

            <SidebarItem name="Klasse" to="/class">
              <PersonIcon></PersonIcon>
            </SidebarItem>

            <SidebarItem name="Hausaufgaben" to="/homework">
              <FileIcon></FileIcon>
            </SidebarItem>

            <SidebarItem name="Einstellungen" to="/settings">
              <GearIcon></GearIcon>
            </SidebarItem>

            <SidebarItem name="Ausloggen" to="/logout">
              <ExitIcon></ExitIcon>
            </SidebarItem>
          </div>
        </SheetContent>
      </Sheet>
    </nav>

    <div class="flex flex-col gap-3 p-3">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import SidebarItem from "~/components/sidebar/SidebarItem.vue";
import { HamburgerMenuIcon, HomeIcon, CircleBackslashIcon, CalendarIcon, PersonIcon, FileIcon, GearIcon, ExitIcon } from "@radix-icons/vue";

const pageName = useRoute().path.replace("/", "").toUpperCase()
</script>