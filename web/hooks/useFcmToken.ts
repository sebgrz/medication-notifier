import firebaseApp from "@/config/firebaseApp";
import CookieManager from "@/managers/cookieManager";
import { getMessaging, getToken, onMessage } from "firebase/messaging";
import { useState } from "react"
import { useEffectOnce } from "./useEffectOnce";

export const useFcmToken = () => {
  const [message, setMessage] = useState("");

  useEffectOnce(() => {
    (async () => {
      const messaging = getMessaging(firebaseApp);

      // Retrieve the notification permission status
      const permission = await Notification.requestPermission();

      // Check if permission is granted before retrieving the token
      if (permission === 'granted') {
        try {
          const currentToken = await getToken(messaging, {
            vapidKey: process.env.NEXT_PUBLIC_FIREBASE_VAPID_KEY,
          });
          if (currentToken) {
            CookieManager.set("fcm-token", currentToken, 5 * 365);
          } else {
            console.error("No registration token available. Request permission to generate one.");
          }

          onMessage(messaging, (payload) => {
            console.info(`msg: ${payload}`);
            setMessage(`${payload}`);
          });
        } catch (e) {
          console.error(`get_token failed: ${e}`);
        }
      }
    })()
  });

  return { message }
}
