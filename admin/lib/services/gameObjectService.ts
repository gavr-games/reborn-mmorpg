import { parseJSONResponse } from "@/lib/utils";

export const getGameObject = (id: string) => {
  return fetch(`/admin/api/game_objects/${id}`).then(parseJSONResponse)
}

export const updateGameObject = (id: string, data: object) => {
  return fetch(`/admin/api/game_objects/${id}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    cache: "no-cache",
    body: JSON.stringify(data)
  }).then(parseJSONResponse)
}

export const deleteGameObject = (id: string) => {
  return fetch(`/admin/api/game_objects/${id}`, {
    method: "DELETE",
  }).then(parseJSONResponse)
}
