import "./index.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

const queryClient = new QueryClient();

const container = document.getElementById("root");
if (container) {
  createRoot(container).render(
    <QueryClientProvider client={queryClient}>
      <StrictMode>
        <div>
          Hello
        </div>
      </StrictMode>
    </QueryClientProvider>,
  );
}
