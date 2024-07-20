import { ChangeEvent } from "react";

const EnumSelector = <TEnum extends object,>(props: { enumType: TEnum, onChange: (v: TEnum) => void }) => {
	const onChange = (event: ChangeEvent<HTMLSelectElement>) => {
		props.onChange((event.target.value as any) as TEnum);
	}

	return (
		<select onChange={onChange}>
			{
				Object.entries(props.enumType).map(([k, _]) => <option key={k} value={k}>{k}</option>)
			}
		</select>
	);
}

export default EnumSelector;
