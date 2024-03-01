import "@/css/style.css"

import Footer from "@/components/Fotter";
import Header from "@/components/Header/Header";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="zh-CN">
      <head>
        <title>我的 WAAA - 论坛</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </head>
      <body>
        <Header />
        <div className="container">
          {children}
        </div>
        <Footer />
      </body>
    </html>
  );
}
