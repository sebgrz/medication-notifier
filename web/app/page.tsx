'use client';

import MedicationsPanel, { Day, Medication, TimeOfDay } from "@/components/medicationsPanel";
import { useState } from "react";

const Index = () => {
  const [medications, setMedications] = useState<Medication[]>([
    { id: "id-123", name: "apap", day: Day.MONDAY, time_of_day: TimeOfDay.MORNING, user_id: "123" },
    { id: "id-124", name: "ibuprofen", day: Day.SATURDAY, time_of_day: TimeOfDay.MIDDAY, user_id: "123" },
  ]);

  return (
    <main className="flex">
      Medications:
      <MedicationsPanel data={medications} />
    </main>
  );
}

export default Index;
