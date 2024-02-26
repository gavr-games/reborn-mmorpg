import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"

export default function ServerControls() {
  return (
    <main className="flex flex-col">
      <h1>Server Controls</h1>
      <Card className="w-full sm:w-64">
        <CardHeader>
          <CardTitle>Engine</CardTitle>
          <CardDescription>Game Engine Service</CardDescription>
        </CardHeader>
        <CardContent>
          Current state: <Badge variant="destructive" className="animate-pulse">Stopped</Badge>
        </CardContent>
        <CardFooter>
          <Button className="bg-green-700">Start</Button>
        </CardFooter>
      </Card>
    </main>
  );
}
