"use client"

import { useState, useEffect } from "react"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import { Button } from "@/components/ui/button"
import { ReloadIcon, Pencil1Icon, TrashIcon } from "@radix-ui/react-icons"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { useToast } from "@/components/ui/use-toast"

import { getGameObject, updateGameObject, deleteGameObject } from "@/lib/services/gameObjectService"

import "jsoneditor/dist/jsoneditor.min.css"

let JSONEditor:any // this trick is required because JSONEditor requires WEB API and produces 500 when Next tries to preload page

const formSchema = z.object({
  id: z.string().min(2, {
    message: "Id must be at least 2 characters.",
  }),
})

export default function GameObjects() {
  const [editor, setEditor] = useState<any>()
  const [gameObjectId, setGameObjectId] = useState<string>()
  const [gameObjectLoading, setGameObjectLoading] = useState<boolean>(false)
  const [gameObjectLoaded, setGameObjectLoaded] = useState<boolean>(false)
  const { toast } = useToast()

  useEffect(() => {
    const initTerminal = async () => {
      JSONEditor = (await import('jsoneditor')).default
    }
    initTerminal()
  }, [])

  // Define form
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: "",
    },
  })

  const handleError = (text:string) => {
    toast({
      variant: "destructive",
      title: "Error",
      description: text,
    })
    if (editor) {
      editor.set("")
    }
    setGameObjectLoading(false)
    setGameObjectLoaded(false)
  }

  const saveGameObject = () => {
    setGameObjectLoading(true)
    updateGameObject(gameObjectId!, editor.get()).then((response) => {
      if (response.ok && response.status === 200) {
        toast({
          title: "Success",
          description: "Game Object updated.",
        })
      } else {
        handleError("Game Object not found")
      }
      setGameObjectLoading(false)
    })
    .catch((error) => {
      handleError(error.message)
    })
  }

  const removeGameObject = () => {
    if (confirm("Do you really want to remove this object?") === true) {
      setGameObjectLoading(true)
      deleteGameObject(gameObjectId!).then((response) => {
        if (response.ok && response.status === 200) {
          toast({
            title: "Success",
            description: "Game Object deleted.",
          })
          if (editor) {
            editor.set("")
          }
        } else {
          handleError("Game Object not found")
        }
        setGameObjectLoading(false)
        setGameObjectLoaded(false)
      })
      .catch((error) => {
        handleError(error.message)
      })
    }
  }

  // Define a submit handler
  function onSubmit(values: z.infer<typeof formSchema>) {
    setGameObjectLoading(true)
    setGameObjectLoaded(false)
    setGameObjectId(values.id)
    getGameObject(values.id)
      .then((response) => {
        if (response.ok && response.status === 200) {
          if (!editor) {
            const container = document.getElementById("json-editor")!
            const options = {}
            const e = new JSONEditor(container, options)
            e.set(response.data)
            setEditor(e)
          } else {
            editor.set(response.data)
          }
          setGameObjectLoaded(true)
        } else {
          handleError("Game Object not found")
        }
        setGameObjectLoading(false)
      })
      .catch((error) => {
        handleError(error.message)
      })
  }

  return (
    <main className="flex flex-col">
      <h1>Find Game Object</h1>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="id"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Id</FormLabel>
                <FormControl>
                  <Input placeholder="3bbceeb1-15c6-4900-87a7-0c51bf7c1603" {...field} />
                </FormControl>
                <FormDescription>
                  This is the ID of Game Object in Redis.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit" disabled={gameObjectLoading}>
          {gameObjectLoading &&
            <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />
          }
            Find
          </Button>
        </form>
      </Form>
      <div className="py-4">
        {gameObjectLoaded &&
          <div className="mb-2">
            <Button onClick={saveGameObject} disabled={gameObjectLoading}>
              <Pencil1Icon className="mr-2 h-4 w-4" />
              Save
            </Button>
            <Button onClick={removeGameObject} variant="destructive" className="ml-2" disabled={gameObjectLoading}>
              <TrashIcon className="mr-2 h-4 w-4" />
              Delete
            </Button>
          </div>
        }
        <div id="json-editor"></div>
      </div>
    </main>
  );
}
