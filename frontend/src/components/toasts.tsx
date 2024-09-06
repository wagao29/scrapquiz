import { CheckCircle2, CircleX } from "lucide-react";
import { toast } from "sonner";

export function toastSuccess(message: string) {
  toast(message, {
    icon: <CheckCircle2 color="green" />,
    position: "top-right",
  });
}

export function toastError(message: string) {
  toast(message, {
    icon: <CircleX color="red" />,
    position: "top-right",
  });
}
