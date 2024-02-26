import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

interface ParsedJSONResponse {
  status: number,
  ok: boolean,
  data: object,
}


export const parseJSONResponse = (response:Response):Promise<ParsedJSONResponse> => {
  return new Promise((resolve) => response.json()
      .then((json) => resolve({
        status: response.status,
        ok: response.ok,
        data: json,
      })));
}
