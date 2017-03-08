using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Drawing;
using System.Drawing.Imaging;
using System.Windows.Forms;
using System.Timers;


class Program
{
    private static System.Timers.Timer aTimer;

    static void get_bmpScreenshot(int i)
    {
        // Creat a new bitmap
        var bmpScreenshot = new Bitmap(Screen.PrimaryScreen.Bounds.Width,
                                       Screen.PrimaryScreen.Bounds.Height,
                                       PixelFormat.Format32bppArgb);
        // Creat a graphic object from the bitmap
        var gfxScreenshot = Graphics.FromImage(bmpScreenshot);

        // Take the screenshot from the upper left corner to the right bottom corner.
        gfxScreenshot.CopyFromScreen(Screen.PrimaryScreen.Bounds.X,
                                    Screen.PrimaryScreen.Bounds.Y,
                                    0,
                                    0,
                                    Screen.PrimaryScreen.Bounds.Size,
                                    CopyPixelOperation.SourceCopy);

        // Save the screenshot to the specified path that the user has chosen.
        var filename = "screenshot" + i.ToString() + ".png";
        bmpScreenshot.Save(filename, ImageFormat.Png);
    }
    static void Main(string[] args)
    {
        SetTimer();
        for (int i = 0; i < 30; i++)
        {
            get_bmpScreenshot(i);
        }
        Console.ReadLine();
        aTimer.Stop();
        aTimer.Dispose();
    }

    private static void SetTimer()
    {
        // Create a timer with a two second interval.
        aTimer = new System.Timers.Timer(2000);
        // Hook up the Elapsed event for the timer. 
        aTimer.Elapsed += OnTimedEvent;
        aTimer.AutoReset = true;
        aTimer.Enabled = true;
    }
    private static void OnTimedEvent(Object source, ElapsedEventArgs e)
    {
        Console.WriteLine("The Elapsed event was raised at {0:HH:mm:ss.fff}",
                          e.SignalTime);
    }
}

