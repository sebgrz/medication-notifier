'use client';

import MedicationsPanel, { Medication } from "@/components/medicationsPanel";
import { useApiManager } from "@/hooks/useApiManager";
import { useCallback, useEffect, useState } from "react";

const Index = () => {
  const api = useApiManager();
  const [medications, setMedications] = useState<Medication[]>([]);

  const loadMedications = useCallback(async () => {
    const freshMedications = await api.fetchMedications();
    setMedications(freshMedications);
  }, []);

  useEffect(() => {
    loadMedications();
  }, [loadMedications]);


  return (
    <main className="flex">
      Medications:
      <MedicationsPanel data={medications} />
    </main>
  );
}

export default Index;
