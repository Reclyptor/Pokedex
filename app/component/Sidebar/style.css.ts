import { createTheme, createThemeContract, style } from "@vanilla-extract/css";

export const vars = createThemeContract({
});

export const theme = createTheme(vars, {
});

export const sidebar = style({
  width: '256px',
  height: '100%',
});