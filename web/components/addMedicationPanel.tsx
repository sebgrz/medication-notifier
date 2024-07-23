'use client';

import { useFormik } from "formik";
import EnumSelector from "./enumSelector";
import { Day, Medication } from "./medicationsPanel";
import { useApiManager } from "@/hooks/useApiManager";
import { useState } from "react";

const AddMedicationPanel = (props: { addedMedicationsAction: (m: Medication[]) => void }) => {
	const api = useApiManager();
	const [day, setDay] = useState(Day.MONDAY);
	const formik = useFormik({
		initialValues: {
			MOR: '',
			MID: '',
			EVE: ''
		},
		onSubmit: async (values) => {
			const medications: Medication[] = [];
			const entries = Object.entries(values).filter(([_, v]) => v.trim().length > 0)
			for (const [k, v] of entries) {
				const medication = await api.appAddMedication(v, day as Day, k);
				if (medication) {
					medications.push(medication);
				}
			}
			props.addedMedicationsAction(medications);
			formik.resetForm();
		},
	});

	const addMedicationSelectDayOnChange = (day: Day) => {
		setDay(day);
	}

	return (
		<tr>
			<th>
				<button type="submit" onClick={() => formik.submitForm()}>ADD</button>&nbsp;
				<EnumSelector enumType={Day} onChange={addMedicationSelectDayOnChange} />
			</th>
			<th><input type="text" id="MOR" name="MOR" onChange={formik.handleChange} value={formik.values.MOR} /></th>
			<th><input type="text" id="MID" name="MID" onChange={formik.handleChange} value={formik.values.MID} /></th>
			<th><input type="text" id="EVE" name="EVE" onChange={formik.handleChange} value={formik.values.EVE} /></th>
		</tr>
	);
}

export default AddMedicationPanel;
