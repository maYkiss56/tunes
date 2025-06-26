import { lazy, Suspense } from "react";
import { createBrowserRouter } from "react-router-dom";
import App from "./App";

import { ProtectedRoute } from "./protectRoutes";

const TracksPage = lazy(() => import("./pages/TracksPage.tsx"));
const Home = lazy(() => import("./pages/Home"));
const NotFoundPage = lazy(() => import("./pages/NotFoundPage"));
const LoginPage = lazy(() => import("./pages/LoginPage"));
const RegisterPage = lazy(() => import("./pages/RegisterPage"));
const ProfilePage = lazy(() => import("./pages/ProfilePage"));
const RatingPage = lazy(() => import("./pages/RatingPage.tsx"));
const AdminAlbumsPage = lazy(() => import("./pages/AdminAlbumsPage"));
const AdminArtistsPage = lazy(() => import("./pages/AdminArtistsPage"));
const AdminGenresPage = lazy(() => import("./pages/AdminGenresPage"));
const AdminPanel = lazy(() => import("./pages/AdminPanel"));
const AdminSongsPage = lazy(() => import("./pages/AdminSongsPage"));

const withSuspense = (element: React.ReactNode) => (
  <Suspense
    fallback={<div className="text-center text-white">Загрузка...</div>}
  >
    {element}
  </Suspense>
);

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <NotFoundPage />,
    children: [
      {
        index: true,
        element: withSuspense(<Home />),
      },
      {
        path: "login",
        element: withSuspense(<LoginPage />),
      },
      {
        path: "register",
        element: withSuspense(<RegisterPage />),
      },
      {
        path: "profile",
        element: withSuspense(
          <ProtectedRoute>
            <ProfilePage />
          </ProtectedRoute>,
        ),
      },
      {
        path: "tracks",
        element: withSuspense(<TracksPage />),
      },
      {
        path: "top",
        element: withSuspense(<RatingPage />),
      },
      {
        path: "admin",
        element: withSuspense(
          <ProtectedRoute adminOnly>
            <AdminPanel />
          </ProtectedRoute>,
        ),
        children: [
          {
            index: true,
            element: <App />,
          },
          {
            path: "songs",
            element: <AdminSongsPage />,
          },
          {
            path: "songs/:id",
            element: <AdminSongsPage />,
          },
          {
            path: "artists",
            element: <AdminArtistsPage />,
          },
          {
            path: "artists/:id",
            element: <AdminArtistsPage />,
          },
          {
            path: "albums",
            element: <AdminAlbumsPage />,
          },
          {
            path: "albums/:id",
            element: <AdminAlbumsPage />,
          },
          {
            path: "genres",
            element: <AdminGenresPage />,
          },
          {
            path: "genres/:id",
            element: <AdminGenresPage />,
          },
        ],
      },
    ],
  },
]);
