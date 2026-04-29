---
name: frontend-architecture
description: Frontend architecture and implementation conventions for app structure, routing, layouts, state ownership, API layers, UI composition, and testing.
---

# Frontend Architecture Reference

## Architecture

```text
main.tsx -> AppProviders -> GlobalRouter -> GlobalLayout -> ModuleRoute
                                                      -> ModuleLayout? -> ModulePage
```

- SPA-only for this phase; no SSR or RSC.
- `app/` owns bootstrap, providers, route composition, and top-level layouts.
- `modules/` own feature routes, pages, queries, API code, schemas, and module-specific UI.
- `shared/` contains only code reused by at least two modules or required globally.

## Module Structure

```text
web/src/
├── app/
│   ├── main/           # App entry
│   ├── providers/      # Query client, router, global hosts
│   ├── layouts/        # Global layouts
│   └── routes/         # Top-level routes, not-found, error boundary
├── modules/
│   └── [module]/
│       ├── router.tsx  # Module route entry
│       ├── layouts/    # Optional module-owned layout
│       ├── pages/      # Route pages
│       ├── components/ # Module-specific UI
│       ├── hooks/      # Module hooks
│       ├── queries/    # Query keys, queries, mutations
│       ├── api/        # Module HTTP DTOs and requests
│       ├── schemas/    # Zod schemas
│       └── types/      # Module types
├── shared/
│   ├── api/            # Base HTTP client and helpers
│   ├── components/     # Reusable composed components
│   ├── hooks/          # Shared hooks
│   ├── lib/            # Framework glue
│   ├── ui/             # shadcn primitives
│   ├── utils/          # Shared utilities
│   ├── config/         # Runtime config
│   └── types/          # Shared types
└── test/               # Setup, mocks, factories
```

- Modules must not import from other modules' internals.
- If reuse is needed across modules, move it into `shared/`.
- `shared/ui/` is only for primitive UI building blocks; product-specific compositions live in `shared/components/` or the module.

## Routing and Layouts

- Use React Router route objects and compose the full tree in `app/routes`.
- Only `app/routes` may mount top-level routes.
- Each module exports a single `router.tsx`; pages never register themselves directly into the global router.
- `GlobalLayout` owns app-wide chrome only: shell, sidebar, header, breadcrumbs container, toasts, dialogs host.
- `ModuleLayout` is optional and module-owned for tabs, filters, sub-navigation, and page scaffolding.
- Modules mount under stable top-level paths such as `/graveyards` or `/people`.
- Use nested routing inside modules for detail, tab, create, and edit flows.

## Server State and Local State

- TanStack Query is the standard for remote fetching, caching, mutations, and invalidation.
- Each module owns its query keys, query functions, and mutation hooks in `queries/`.
- Centralize query keys per module in `queries/keys.ts`.
- URL state stores filters, pagination, sorting, and tab selection when it should survive navigation or be shareable.
- Local component state is only for ephemeral UI state such as modal visibility, hover state, and temporary drafts.
- React Context is reserved for cross-cutting concerns such as theme, locale, auth session, or feature flags.
- Redux and Zustand are not used.

## API and Contracts

- Use `shared/api/http.ts` for the base HTTP client.
- Use `shared/api/client.ts` for common request helpers, headers, and response normalization.
- Keep transport DTOs handwritten and colocated in `modules/[module]/api/` unless shared across modules.
- Do not call raw `fetch()` from pages or components.
- Modules should consume API helpers or query hooks, not transport details directly.

## UI and Forms

- Use shadcn/ui for primitives; generated components live in `shared/ui/`.
- Define and reuse design tokens for colors, spacing, radius, typography, shadows, and semantic states.
- Keep the set of global layouts small; use `BlankLayout` only for materially different full-screen flows.
- Default form stack: `react-hook-form` + `zod` + shadcn form primitives.
- Standardize loading, empty, error, and success states.
- Provide a single app-level toast system.
- Mutation UX defaults: disable submit while pending, show inline validation errors, reserve toasts for cross-page or high-value feedback.

## Error Handling and Providers

- `AppProviders` composes `QueryClientProvider`, `RouterProvider`, theme provider if needed, toaster, and dialog hosts.
- Add route-level error boundaries and a shared not-found route.
- Normalize API errors in shared code so modules do not parse errors differently.
- Centralize runtime config in `shared/config/`; do not read environment variables directly inside modules.
